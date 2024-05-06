package quote

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-resty/resty/v2"

	"quote/config"
	"quote/pkg/database"
)

const (
	randomQuoteURL            = "https://api.quotable.io/random"
	oneHundredPercent float64 = 100
)

type Service struct {
	cfg   *config.QuotesConfig
	db    Database
	resty *resty.Client
}

func (s *Service) GetQuote(ctx context.Context, userID string) (Quote, error) {
	quotes, err := s.db.GetQuotes(ctx, userID)
	if err != nil {
		return Quote{}, fmt.Errorf("failed to get quotes: %w", err)
	}

	quote, err := s.getQuote(ctx, quotes)
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
	quote, err := s.db.GetSameQuote(ctx, userID, quoteID)
	if err != nil && !errors.Is(err, database.ErrRecordNotFound) {
		return Quote{}, fmt.Errorf("failed to get same quote: %w", err)
	} else if errors.Is(err, database.ErrRecordNotFound) {
		quote, err = s.getRandomQuote(ctx)
		if err != nil {
			return Quote{}, fmt.Errorf("failed to get random quote: %w", err)
		}
	}

	if err = s.db.MarkAsViewed(ctx, userID, quote.ID); err != nil {
		return Quote{}, fmt.Errorf("failed to mark as viewed: %w", err)
	}

	return fromDatabaseQuoteToQuote(quote), nil
}

func NewService(cfg *config.Config, db Database, resty *resty.Client) *Service {
	return &Service{
		cfg:   &cfg.QuotesConfig,
		db:    db,
		resty: resty,
	}
}

func (s *Service) getQuote(ctx context.Context, quotes []database.Quote) (database.Quote, error) {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	randomPercent := rnd.Float64() * oneHundredPercent

	if (oneHundredPercent-s.cfg.RandomQuoteChance) > randomPercent && len(quotes) > 0 {
		var likes float64
		for _, q := range quotes {
			if q.Likes == 0 {
				q.Likes++
			}

			likes += float64(q.Likes)
		}

		var accumulator float64
		del := likes * oneHundredPercent / (oneHundredPercent - s.cfg.RandomQuoteChance)
		for i, q := range quotes {
			if q.Likes == 0 {
				q.Likes++
			}

			percent := float64(q.Likes) / del * oneHundredPercent
			if percent+accumulator >= randomPercent {
				return quotes[i], nil
			}

			accumulator += percent
		}
	}

	return s.getRandomQuote(ctx)
}

func (s *Service) getRandomQuote(ctx context.Context) (database.Quote, error) {
	resp, err := s.resty.R().Get(randomQuoteURL)
	if err != nil {
		return database.Quote{}, fmt.Errorf("failed to receive random quote from site: %w", err)
	}

	var randomQuote RandomQuote
	err = json.Unmarshal(resp.Body(), &randomQuote)
	if err != nil {
		return database.Quote{}, fmt.Errorf("failed to unmarshal random quote: %w", err)
	}

	quote := randomQuote.toDatabase()
	err = s.db.SaveQuote(ctx, quote)
	if err != nil {
		return database.Quote{}, fmt.Errorf("failed to save new random quote: %w", err)
	}

	return quote, nil
}
