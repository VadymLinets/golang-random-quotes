package main

import (
	"go.uber.org/fx"

	"quote/cmd"
	"quote/config"
)

//go:generate go get -d github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	fx.New(cmd.Exec(cfg)).Run()
}
