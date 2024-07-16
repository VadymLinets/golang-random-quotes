package handlers

import (
	"context"
	"log/slog"

	pb "quote/server/proto"
)

func (h *Handler) GetQuoteHandler(ctx context.Context, in *pb.UserIDRequest) (*pb.Quote, error) {
	quote, err := h.quotes.GetQuote(ctx, in.GetUserId())
	if err != nil {
		slog.Error("Failed to get quote:", "err", err)
		return nil, err
	}

	return &pb.Quote{
		Id:     quote.ID,
		Quote:  quote.Quote,
		Author: quote.Author,
		Tags:   quote.Tags,
		Likes:  quote.Likes,
	}, nil
}

func (h *Handler) GetSameQuoteHandler(ctx context.Context, in *pb.UserAndQuoteIDRequest) (*pb.Quote, error) {
	quote, err := h.quotes.GetSameQuote(ctx, in.GetUserId(), in.GetQuoteId())
	if err != nil {
		slog.Error("Failed to get same quote:", "quote id", in.GetQuoteId(), "err", err)
		return nil, err
	}

	return &pb.Quote{
		Id:     quote.ID,
		Quote:  quote.Quote,
		Author: quote.Author,
		Tags:   quote.Tags,
		Likes:  quote.Likes,
	}, nil
}

func (h *Handler) LikeQuoteHandler(ctx context.Context, in *pb.UserAndQuoteIDRequest) (*pb.Empty, error) {
	err := h.quotes.LikeQuote(ctx, in.GetUserId(), in.GetQuoteId())
	if err != nil {
		slog.Error("Failed to like quote:", "quote id", in.GetQuoteId(), "err", err)
		return nil, err
	}

	return &pb.Empty{}, nil
}
