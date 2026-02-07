package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SystemRepository interface {
	Ping(ctx context.Context) error
}

type systemRepository struct {
	pool *pgxpool.Pool
}

func NewSystemRepository(pool *pgxpool.Pool) SystemRepository {
	return &systemRepository{pool: pool}
}

func (r *systemRepository) Ping(ctx context.Context) error {
	return r.pool.Ping(ctx)
}
