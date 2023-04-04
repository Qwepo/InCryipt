package postgres

import (
	"context"
	"fmt"

	internal "github.com/Qwepo/InCryipt/Internal"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(ctx context.Context, conf *internal.Config) (*pgxpool.Conn, error) {
	// url = postgresql://username:password@127.0.0.1:8080/dbname
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.Address, conf.Postgres.Port, conf.Postgres.Database)
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	if pool.Ping(ctx) != nil {
		return nil, err
	}
	return conn, err
}
