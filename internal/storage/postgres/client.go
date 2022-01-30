package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"mycalendar/config"
)

func NewClient(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName)

	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL")
	}

	return pool
}
