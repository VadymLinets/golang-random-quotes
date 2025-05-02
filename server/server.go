package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	"quote/config"
	"quote/server/graphql"
	"quote/server/handlers"
	"quote/server/middlewares"
)

type HTTPServer struct {
	server *http.Server

	cfg      *config.ServerConfig
	handlers *handlers.Handler
	resolver *graphql.Resolver
}

func (s *HTTPServer) Start(_ context.Context) error {
	gin.SetMode(gin.ReleaseMode)

	srv := handler.New(graphql.NewExecutableSchema(graphql.Config{Resolvers: s.resolver}))

	router := gin.New()
	router.Use(
		middlewares.CorsMiddleware(s.cfg.CorsMaxAge),
		middlewares.LoggerMiddleware(),
		gin.Recovery(),
	)

	// System
	router.GET("/heartbeat", s.handlers.Heartbeat)

	// Quotes
	router.GET("/", s.handlers.GetQuoteHandler)
	router.GET("/same", s.handlers.GetSameQuoteHandler)
	router.PATCH("/like", s.handlers.LikeQuoteHandler)

	// GraphQL
	router.POST("/graphql", gin.WrapH(srv))

	s.server = &http.Server{
		Addr:              s.cfg.Addr,
		Handler:           router,
		ReadHeaderTimeout: s.cfg.ReadHeaderTimeout,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server start error", "err", err)
			panic(err)
		}
	}()

	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func NewHTTPServer(cfg *config.Config, handlers *handlers.Handler, resolver *graphql.Resolver) *HTTPServer {
	return &HTTPServer{
		cfg:      &cfg.ServerConfig,
		handlers: handlers,
		resolver: resolver,
	}
}
