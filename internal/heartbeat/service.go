package heartbeat

import (
	"context"
	"log/slog"
)

type Service struct {
	db Database
}

func (s *Service) Ping(ctx context.Context) error {
	err := s.db.Ping(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to ping database", "err", err)
		return err
	}

	return nil
}

func NewService(db Database) *Service {
	return &Service{db: db}
}
