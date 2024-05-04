package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"

	"quote/config"
)

type Gorm struct {
	db  *gorm.DB
	cfg *config.GormConfig
}

func (p *Gorm) Ping(ctx context.Context) error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}

	return db.PingContext(ctx)
}

func (p *Gorm) GetQuote(ctx context.Context, quoteID string) (quote Quote, err error) {
	err = p.db.Table("quotes").
		WithContext(ctx).
		Select([]string{
			"id",
			"quote",
			"author",
			"likes",
		}).
		Where("quotes.id = ?", quoteID).
		First(&quote).Error
	return
}

func (p *Gorm) GetQuotes(ctx context.Context, userID string) (quotes []Quote, err error) {
	viewed := p.db.Table("quotes").
		WithContext(ctx).
		Select("quotes.id").
		Joins("INNER JOIN views ON views.quote_id = quotes.id").
		Where("views.user_id = ?", userID)

	err = p.db.Table("quotes").
		WithContext(ctx).
		Select([]string{
			"id",
			"quote",
			"author",
			"likes",
		}).
		Where("quotes.id NOT IN (?)", viewed).
		Order("quotes.likes DESC").
		Find(&quotes).Error
	return
}

func (p *Gorm) GetSameQuote(ctx context.Context, userID, quoteID string) (quote Quote, err error) {
	viewed := p.db.Table("quotes").
		Select("quotes.id").
		Joins("INNER JOIN views ON views.quote_id = quotes.id").
		Where("views.user_id = ?", userID)

	author := fmt.Sprintf(`SELECT author FROM quotes WHERE id = '%s'`, quoteID)

	err = p.db.Table("quotes").
		WithContext(ctx).
		Select([]string{
			"id",
			"quote",
			"author",
			"likes",
		}).
		Where("quotes.id NOT IN (?)", viewed).
		Order("(case when quotes.author = (" + author + ") then 1 else 2 end)").
		Order("quotes.likes DESC").
		Limit(1).
		First(&quote).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrRecordNotFound
	}
	return
}

func (p *Gorm) GetView(ctx context.Context, userID, quoteID string) (view View, err error) {
	err = p.db.Table("views").
		WithContext(ctx).
		Where("user_id = ? AND quote_id = ?", userID, quoteID).
		First(&view).Error
	return
}

func (p *Gorm) SaveQuote(ctx context.Context, quote Quote) error {
	return p.db.Table("quotes").WithContext(ctx).Save(quote).Error
}

func (p *Gorm) MarkAsViewed(ctx context.Context, userID, quoteID string) error {
	return p.db.Table("views").WithContext(ctx).Create(View{
		UserID:  userID,
		QuoteID: quoteID,
	}).Error
}

func (p *Gorm) MarkAsLiked(ctx context.Context, userID, quoteID string) error {
	return p.db.Table("views").
		WithContext(ctx).
		Where("user_id = ? AND quote_id = ?", userID, quoteID).
		Update("liked", true).Error
}

func (p *Gorm) LikeQuote(ctx context.Context, quoteID string) error {
	return p.db.Table("quotes").
		WithContext(ctx).
		Where("id = ?", quoteID).
		UpdateColumn("likes", gorm.Expr("likes + 1")).Error
}

func (p *Gorm) Start(_ context.Context) (err error) {
	var dialector gorm.Dialector

	switch p.cfg.Dialect {
	case "mysql":
		dialector = mysql.Open(p.cfg.DSN)
	case "postgres":
		dialector = postgres.Open(p.cfg.DSN)
	}

	p.db, err = gorm.Open(dialector, &gorm.Config{
		Logger: glog.Default.LogMode(glog.Silent),
	})
	if err != nil {
		return
	}

	db, err := p.db.DB()
	if err != nil {
		return
	}

	if err = goose.SetDialect(p.cfg.Dialect); err != nil {
		return
	}

	err = goose.Up(db, p.cfg.MigrationPath, goose.WithAllowMissing())
	return
}

func (p *Gorm) Stop(ctx context.Context) error {
	sql, err := p.db.WithContext(ctx).DB()
	if err != nil {
		return err
	}

	err = sql.Close()
	if err != nil {
		return err
	}

	return nil
}

func NewGorm(cfg *config.Config) *Gorm {
	return &Gorm{cfg: &cfg.GormConfig}
}
