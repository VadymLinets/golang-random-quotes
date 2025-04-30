package main

import (
	"go.uber.org/fx"

	"quote/app"
	"quote/config"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	fx.New(app.Exec(cfg)).Run()
}
