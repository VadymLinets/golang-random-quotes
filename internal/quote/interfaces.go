package quote

import (
	"context"

	"quote/pkg/database"
)

type Database interface {
	GetQuotes(ctx context.Context, userID string) ([]database.Quote, error)
	GetSameQuote(ctx context.Context, userID, author string) (database.Quote, error)
	GetView(ctx context.Context, userID, quoteID string) (database.View, error)
	SaveQuote(ctx context.Context, quote database.Quote) error
	MarkAsViewed(ctx context.Context, userID, quoteID string) error
	MarkAsLiked(ctx context.Context, userID, quoteID string) error
	LikeQuote(ctx context.Context, quoteID string) error
}
