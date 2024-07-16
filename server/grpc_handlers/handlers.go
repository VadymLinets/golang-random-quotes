package handlers

import (
	"context"
	"log/slog"

	"quote/internal/heartbeat"
	"quote/internal/quote"
	pb "quote/server/proto"
)

type Handler struct {
	pb.UnimplementedQuotesServer

	quotes    *quote.Service
	heartbeat *heartbeat.Service
}

func (h *Handler) Heartbeat(ctx context.Context, _ *pb.Empty) (*pb.Empty, error) {
	err := h.heartbeat.PingDatabase(ctx)
	if err != nil {
		slog.Error("Failed to ping database", "err", err)
		return nil, err
	}

	return &pb.Empty{}, nil
}

func NewHandler(quotes *quote.Service, heartbeat *heartbeat.Service) *Handler {
	return &Handler{
		quotes:    quotes,
		heartbeat: heartbeat,
	}
}
