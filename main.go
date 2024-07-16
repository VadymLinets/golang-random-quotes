package main

import (
	"go.uber.org/fx"

	"quote/cmd"
	"quote/config"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	fx.New(cmd.Exec(cfg)).Run()
}
