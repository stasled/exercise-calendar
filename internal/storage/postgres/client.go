package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func NewClient(ctx context.Context) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_HOST"),
		viper.GetInt("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DB"))

	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL")
	}

	return pool
}
