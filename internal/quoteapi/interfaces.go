package quoteapi

import (
	"context"

	"quote/pkg/database"
)

type Database interface {
	SaveQuote(ctx context.Context, quote database.Quote) error
}
