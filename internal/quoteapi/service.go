package quoteapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"

	"quote/pkg/database"
)

const (
	randomQuoteURL = "https://dummyjson.com/quotes/random"
)

type Service struct {
	db    Database
	resty *resty.Client
}

func (s *Service) GetRandomQuote(ctx context.Context) (database.Quote, error) {
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

func NewService(db Database, resty *resty.Client) *Service {
	return &Service{
		db:    db,
		resty: resty,
	}
}
