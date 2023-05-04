package postgres

import (
	"context"
	"fmt"

	internal "github.com/Qwepo/InCryipt/internal"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
	url            string
	migrationspath string
}

func NewClient(ctx context.Context, conf *internal.Config) (*DB, error) {
	// url = postgresql://username:password@127.0.0.1:8080/dbname?sslmode=false
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.Address, conf.Postgres.Port, conf.Postgres.Database)
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	if pool.Ping(ctx) != nil {
		return nil, err
	}

	return &DB{Pool: pool, url: url, migrationspath: conf.Postgres.MigrationsPaht}, nil
}

func (db *DB) Migrate() error {
	mgpath := fmt.Sprintf("file://%s", db.migrationspath)
	m, err := migrate.New(mgpath, db.url)
	if err != nil {
		return err
	}
	defer m.Close()
	err = m.Up()
	switch err {
	case migrate.ErrNoChange:
		return nil
	default:
		return err
	}

}
