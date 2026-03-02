package main

import (
	"github.com/lunai-monster/lunar-pos/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		server.Module,
	).Run()
}
