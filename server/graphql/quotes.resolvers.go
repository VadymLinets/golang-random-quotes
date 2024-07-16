package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
)

// LikeQuoteHandler is the resolver for the like_quote_handler field.
func (r *mutationHandlerResolver) LikeQuoteHandler(ctx context.Context, userID string, quoteID string) (*EmptyResult, error) {
	err := r.QuotesHandler.LikeQuote(ctx, userID, quoteID)
	if err != nil {
		return &EmptyResult{
			Success: false,
			Errors:  []string{err.Error()},
		}, nil
	}

	return &EmptyResult{Success: true}, nil
}

// Heartbeat is the resolver for the heartbeat field.
func (r *queryHandlerResolver) Heartbeat(ctx context.Context) (*EmptyResult, error) {
	err := r.HeartbeatHandler.PingDatabase(ctx)
	if err != nil {
		return &EmptyResult{
			Success: false,
			Errors:  []string{err.Error()},
		}, nil
	}

	return &EmptyResult{Success: true}, nil
}

// GetQuoteHandler is the resolver for the get_quote_handler field.
func (r *queryHandlerResolver) GetQuoteHandler(ctx context.Context, userID string) (*QuoteResult, error) {
	quote, err := r.QuotesHandler.GetQuote(ctx, userID)
	if err != nil {
		return &QuoteResult{
			Success: false,
			Errors:  []string{err.Error()},
		}, nil
	}

	return &QuoteResult{Success: true, Quote: &Quote{
		ID:     quote.ID,
		Quote:  quote.Quote,
		Author: quote.Author,
		Tags:   quote.Tags,
		Likes:  int(quote.Likes),
	}}, nil
}

// GetSameQuoteHandler is the resolver for the get_same_quote_handler field.
func (r *queryHandlerResolver) GetSameQuoteHandler(ctx context.Context, userID string, quoteID string) (*QuoteResult, error) {
	quote, err := r.QuotesHandler.GetSameQuote(ctx, userID, quoteID)
	if err != nil {
		return &QuoteResult{
			Success: false,
			Errors:  []string{err.Error()},
		}, nil
	}

	return &QuoteResult{Success: true, Quote: &Quote{
		ID:     quote.ID,
		Quote:  quote.Quote,
		Author: quote.Author,
		Tags:   quote.Tags,
		Likes:  int(quote.Likes),
	}}, nil
}

// MutationHandler returns MutationHandlerResolver implementation.
func (r *Resolver) MutationHandler() MutationHandlerResolver { return &mutationHandlerResolver{r} }

// QueryHandler returns QueryHandlerResolver implementation.
func (r *Resolver) QueryHandler() QueryHandlerResolver { return &queryHandlerResolver{r} }

type (
	mutationHandlerResolver struct{ *Resolver }
	queryHandlerResolver    struct{ *Resolver }
)
