package handlers

import (
	"context"

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
	err := h.heartbeat.Ping(ctx)
	if err != nil {
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
