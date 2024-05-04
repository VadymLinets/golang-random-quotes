package cmd

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"

	"quote/config"
	"quote/internal/heartbeat"
	"quote/internal/quote"
	"quote/pkg/database"
	"quote/server"
	"quote/server/handlers"
)

func Exec(cfg *config.Config) fx.Option {
	return fx.Options(
		fx.Provide(
			func() *config.Config { return cfg },
			database.NewGorm,
			fx.Annotate(
				copyForAnnotation[database.Gorm],
				fx.As(new(heartbeat.Database)),
				fx.As(new(quote.Database)),
			),
			resty.New,
			quote.NewService,
			heartbeat.NewService,
			handlers.NewHandler,
			server.NewHTTPServer,
		),
		fx.Invoke(
			prepareHooks,
		),
	)
}

func copyForAnnotation[T any](v *T) *T {
	return v
}

type hooks struct {
	fx.In

	Database *database.Gorm
	Server   *server.HTTPServer
}

func prepareHooks(lc fx.Lifecycle, hooks hooks) {
	lc.Append(fx.Hook{OnStart: hooks.Database.Start, OnStop: hooks.Database.Stop})
	lc.Append(fx.Hook{OnStart: hooks.Server.Start, OnStop: hooks.Server.Stop})
}
