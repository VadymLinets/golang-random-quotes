package cmd

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"quote/config"
	"quote/internal/heartbeat"
	"quote/internal/quote"
	"quote/internal/quoteapi"
	"quote/pkg/database"
	"quote/server"
	"quote/server/graphql"
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
			server.NewHTTPServer,
		),
		fx.Invoke(
			prepareHooks,
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
	Server   *server.HTTPServer
}

func prepareHooks(lc fx.Lifecycle, hooks hooks) {
	lc.Append(fx.Hook{OnStart: hooks.Database.Start, OnStop: hooks.Database.Stop})
	lc.Append(fx.Hook{OnStart: hooks.Server.Start, OnStop: hooks.Server.Stop})
}
