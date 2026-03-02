package database

import (
	"context"
	"database/sql"

	"github.com/lunai-monster/lunar-pos/internal/config"
	"github.com/lunai-monster/lunar-pos/internal/database/sqlc"
	"go.uber.org/fx"
	_ "modernc.org/sqlite"
)

var Module = fx.Options(
	fx.Provide(NewConnection),
	fx.Provide(NewQueries),
)

func NewConnection(lc fx.Lifecycle, cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.DBURL)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.PingContext(ctx); err != nil {
				return err
			}
			schema := `
				CREATE TABLE IF NOT EXISTS products (
			  sku TEXT PRIMARY KEY NOT NULL,
			  title text not null,
			  pricecents integer not null 
			)
			`
			_, err := db.ExecContext(ctx, schema)
			return err
		},

		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})
	return db, nil
}

func NewQueries(conn *sql.DB) *sqlc.Queries {
	return sqlc.New(conn)
}
