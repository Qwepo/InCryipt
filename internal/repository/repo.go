package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	dbUser
	Close(context.Context) error
}

type client struct {
	*pgxpool.Conn
}

func NewRepo(ctx context.Context, db *pgxpool.Pool) (Repository, error) {
	conn, err := db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	return &client{conn}, nil
}

func (c *client) Close(ctx context.Context) error {
	return c.Conn.Conn().Close(ctx)
}
