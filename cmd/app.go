package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"quote/config"
	"quote/internal/heartbeat"
	"quote/internal/quote"
	"quote/internal/quoteapi"
	"quote/pkg/database"
	"quote/server"
	"quote/server/graphql"
	grpcHandlers "quote/server/grpc_handlers"
	"quote/server/handlers"
)

func Exec(cfg *config.Config) fx.Option {
	return fx.Options(
		fx.Provide(
			func() *config.Config { return cfg },
			database.NewPostgres,
			fx.Annotate(
				copyForAnnotation[database.Postgres],
				fx.As(new(heartbeat.Database)),
				fx.As(new(quote.Database)),
				fx.As(new(quoteapi.Database)),
			),
			resty.New,
			fx.Annotate(quoteapi.NewService, fx.As(new(quote.API))),
			quote.NewService,
			heartbeat.NewService,
			newGraphQLResolver,
			handlers.NewHandler,
			grpcHandlers.NewHandler,
		),
		fx.Invoke(
			prepareHooks,
			startServer,
		),
	)
}

func newGraphQLResolver(quotes *quote.Service, heartbeat *heartbeat.Service) *graphql.Resolver {
	return &graphql.Resolver{
		QuotesHandler:    quotes,
		HeartbeatHandler: heartbeat,
	}
}

func copyForAnnotation[T any](v *T) *T {
	return v
}

type hooks struct {
	fx.In

	Database *database.Postgres
}

func prepareHooks(lc fx.Lifecycle, in hooks) {
	lc.Append(fx.Hook{OnStart: in.Database.Start, OnStop: in.Database.Stop})
}

type srv struct {
	fx.In

	Cfg          *config.Config
	Handlers     *handlers.Handler
	GrpcHandlers *grpcHandlers.Handler
	Resolver     *graphql.Resolver
}

func startServer(lc fx.Lifecycle, in srv) error {
	switch in.Cfg.Type {
	case config.HTTP:
		http := server.NewHTTPServer(in.Cfg, in.Handlers, in.Resolver)
		lc.Append(fx.Hook{OnStart: http.Start, OnStop: http.Stop})
	case config.GRPC:
		grpc := server.NewGRPCServer(in.Cfg, in.GrpcHandlers)
		lc.Append(fx.Hook{OnStart: grpc.Start, OnStop: grpc.Stop})
	default:
		return fmt.Errorf("unsupported server type: %s", in.Cfg.Type)
	}

	return nil
}
