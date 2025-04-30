package quote

import (
	"context"
	"errors"
	"fmt"

	"quote/config"
	"quote/pkg/database"
	"quote/tools"
)

type Service struct {
	cfg *config.QuotesConfig
	db  Database
	api API
}

func (s *Service) GetQuote(ctx context.Context, userID string) (Quote, error) {
	quotes, err := s.db.GetQuotes(ctx, userID)
	if err != nil {
		return Quote{}, fmt.Errorf("failed to get quotes: %w", err)
	}

	quote, err := s.chooseQuote(ctx, quotes, tools.RandomPercent(oneHundredPercent))
	if err != nil {
		return Quote{}, fmt.Errorf("failed to get random quote: %w", err)
	}

	if err = s.db.MarkAsViewed(ctx, userID, quote.ID); err != nil {
		return Quote{}, fmt.Errorf("failed to mark as viewed: %w", err)
	}

	return fromDatabaseQuoteToQuote(quote), nil
}

func (s *Service) LikeQuote(ctx context.Context, userID, quoteID string) error {
	view, err := s.db.GetView(ctx, userID, quoteID)
	if err != nil {
		return fmt.Errorf("failed to get view: %w", err)
	}

	if view.Liked {
		return nil
	}

	err = s.db.LikeQuote(ctx, quoteID)
	if err != nil {
		return fmt.Errorf("failed to like quote: %w", err)
	}

	err = s.db.MarkAsLiked(ctx, userID, quoteID)
	if err != nil {
		return fmt.Errorf("failed to mark as liked: %w", err)
	}

	return nil
}

func (s *Service) GetSameQuote(ctx context.Context, userID, quoteID string) (Quote, error) {
	viewedQuote, err := s.db.GetQuote(ctx, quoteID)
	if err != nil {
		return Quote{}, fmt.Errorf("failed to get viewed quote: %w", err)
	}

	quote, err := s.db.GetSameQuote(ctx, userID, viewedQuote)
	if err != nil && !errors.Is(err, database.ErrRecordNotFound) {
		return Quote{}, fmt.Errorf("failed to get same quote: %w", err)
	} else if errors.Is(err, database.ErrRecordNotFound) {
		quote, err = s.api.GetRandomQuote(ctx)
		if err != nil {
			return Quote{}, fmt.Errorf("failed to get random quote: %w", err)
		}
	}

	if err = s.db.MarkAsViewed(ctx, userID, quote.ID); err != nil {
		return Quote{}, fmt.Errorf("failed to mark as viewed: %w", err)
	}

	return fromDatabaseQuoteToQuote(quote), nil
}

func NewService(cfg *config.Config, db Database, api API) *Service {
	return &Service{
		cfg: &cfg.QuotesConfig,
		db:  db,
		api: api,
	}
}

func (s *Service) chooseQuote(ctx context.Context, quotes []database.Quote, randomPercent float64) (database.Quote, error) {
	if len(quotes) == 0 {
		return s.api.GetRandomQuote(ctx)
	}

	if (oneHundredPercent - s.cfg.RandomQuoteChance) > randomPercent {
		likesTotal := tools.Sum(quotes, func(quote database.Quote) float64 {
			if quote.Likes == 0 {
				return 1
			}

			return float64(quote.Likes)
		})

		var accumulator float64
		del := likesTotal * oneHundredPercent / (oneHundredPercent - s.cfg.RandomQuoteChance)
		for _, q := range quotes {
			quoteLikes := 1.0
			if q.Likes != 0 {
				quoteLikes = float64(q.Likes)
			}

			percent := quoteLikes / del * oneHundredPercent
			if percent+accumulator >= randomPercent {
				return q, nil
			}

			accumulator += percent
		}
	}

	return s.api.GetRandomQuote(ctx)
}
