package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"

	"quote/config"
)

type Postgres struct {
	db  *gorm.DB
	cfg *config.PostgresConfig
}

func (p *Postgres) Ping(ctx context.Context) error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}

	return db.PingContext(ctx)
}

func (p *Postgres) GetQuote(ctx context.Context, quoteID string) (quote Quote, err error) {
	err = p.db.Table("quotes").
		WithContext(ctx).
		Select("*").
		Where("quotes.id = ?", quoteID).
		First(&quote).Error
	return
}

func (p *Postgres) GetQuotes(ctx context.Context, userID string) (quotes []Quote, err error) {
	viewed := p.db.Table("quotes").
		WithContext(ctx).
		Select("quotes.id").
		Joins("INNER JOIN views ON views.quote_id = quotes.id").
		Where("views.user_id = ?", userID)

	err = p.db.Table("quotes").
		WithContext(ctx).
		Select("*").
		Where("quotes.id NOT IN (?)", viewed).
		Order("quotes.likes DESC").
		Find(&quotes).Error
	return
}

func (p *Postgres) GetSameQuote(ctx context.Context, userID, quoteID string) (quote Quote, err error) {
	viewed := p.db.Table("quotes").
		Select("quotes.id").
		Joins("INNER JOIN views ON views.quote_id = quotes.id").
		Where("views.user_id = ?", userID)

	author := fmt.Sprintf(`SELECT author FROM quotes WHERE id = '%s'`, quoteID)

	err = p.db.Table("quotes").
		WithContext(ctx).
		Select("*").
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

func (p *Postgres) GetView(ctx context.Context, userID, quoteID string) (view View, err error) {
	err = p.db.Table("views").
		WithContext(ctx).
		Where("user_id = ? AND quote_id = ?", userID, quoteID).
		First(&view).Error
	return
}

func (p *Postgres) SaveQuote(ctx context.Context, quote Quote) error {
	return p.db.Table("quotes").WithContext(ctx).Save(quote).Error
}

func (p *Postgres) MarkAsViewed(ctx context.Context, userID, quoteID string) error {
	return p.db.Table("views").WithContext(ctx).Create(View{
		UserID:  userID,
		QuoteID: quoteID,
	}).Error
}

func (p *Postgres) MarkAsLiked(ctx context.Context, userID, quoteID string) error {
	return p.db.Table("views").
		WithContext(ctx).
		Where("user_id = ? AND quote_id = ?", userID, quoteID).
		Update("liked", true).Error
}

func (p *Postgres) LikeQuote(ctx context.Context, quoteID string) error {
	return p.db.Table("quotes").
		WithContext(ctx).
		Where("id = ?", quoteID).
		UpdateColumn("likes", gorm.Expr("likes + 1")).Error
}

func (p *Postgres) Start(_ context.Context) (err error) {
	p.db, err = gorm.Open(postgres.Open(p.cfg.DSN), &gorm.Config{
		Logger: glog.Default.LogMode(glog.Silent),
	})
	if err != nil {
		return
	}

	db, err := p.db.DB()
	if err != nil {
		return
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return
	}

	err = goose.Up(db, p.cfg.MigrationPath, goose.WithAllowMissing())
	return
}

func (p *Postgres) Stop(ctx context.Context) error {
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

func NewPostgres(cfg *config.Config) *Postgres {
	return &Postgres{cfg: &cfg.PostgresConfig}
}
