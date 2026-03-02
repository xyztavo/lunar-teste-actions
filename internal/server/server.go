package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lunai-monster/lunar-pos/internal/config"
	"github.com/lunai-monster/lunar-pos/internal/database"
	"github.com/lunai-monster/lunar-pos/internal/handlers"
	"github.com/lunai-monster/lunar-pos/internal/routes"
	"github.com/lunai-monster/lunar-pos/internal/utils"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	database.Module,
	utils.Module,
	handlers.Module,
	routes.Module,
	fx.Provide(NewServer),
	fx.Invoke(registerHooks),
)

func NewServer(cfg *config.Config, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}
}

func registerHooks(lc fx.Lifecycle, srv *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Listening at ", srv.Addr)
			go srv.Serve(ln)
			return nil
		},

		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
