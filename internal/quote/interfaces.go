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
	SaveQuote(ctx context.Context, quote database.Quote) error
	MarkAsViewed(ctx context.Context, userID, quoteID string) error
	MarkAsLiked(ctx context.Context, userID, quoteID string) error
	LikeQuote(ctx context.Context, quoteID string) error
}
