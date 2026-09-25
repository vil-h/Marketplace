package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(
		ctx,
		"postgres://localhost:5432/marketplace")
}
