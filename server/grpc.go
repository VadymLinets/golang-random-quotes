package server

import (
	"context"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"quote/config"
	pb "quote/server/proto"
)

type GRPCServer struct {
	cfg      *config.ServerConfig
	listener net.Listener
	server   *grpc.Server

	quotesHandler pb.QuotesServer
}

func (s *GRPCServer) Start(_ context.Context) error {
	var err error
	s.listener, err = net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		return err
	}

	s.server = grpc.NewServer()
	pb.RegisterQuotesServer(s.server, s.quotesHandler)

	reflection.Register(s.server)
	go func() {
		if err := s.server.Serve(s.listener); err != nil {
			slog.Error("Server start error", "err", err)
			panic(err)
		}
	}()

	return nil
}

func (s *GRPCServer) Stop(_ context.Context) error {
	s.server.GracefulStop()
	return s.listener.Close()
}

func NewGRPCServer(cfg *config.Config, quotesHandler pb.QuotesServer) *GRPCServer {
	return &GRPCServer{cfg: &cfg.ServerConfig, quotesHandler: quotesHandler}
}
