package quote

import (
	"context"

	"quote/pkg/database"
)

type Database interface {
	GetQuote(ctx context.Context, quoteID string) (database.Quote, error)
	GetQuotes(ctx context.Context, userID string) ([]database.Quote, error)
	GetSameQuote(ctx context.Context, userID string, viewedQuote database.Quote) (database.Quote, error)
	GetView(ctx context.Context, userID, quoteID string) (database.View, error)
	MarkAsViewed(ctx context.Context, userID, quoteID string) error
	MarkAsLiked(ctx context.Context, userID, quoteID string) error
	LikeQuote(ctx context.Context, quoteID string) error
}

type API interface {
	GetRandomQuote(ctx context.Context) (database.Quote, error)
}
