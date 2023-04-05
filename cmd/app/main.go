package main

import (
	"context"

	internal "github.com/Qwepo/InCryipt/internal"
	"github.com/Qwepo/InCryipt/pkg/database/postgres"
	"github.com/Qwepo/InCryipt/pkg/logger"
)

func main() {
	ctx := context.TODO()
	conf, err := internal.NewConfig()
	if err != nil {
		panic(err)
	}

	// create logger
	log := logger.NewLogger(conf)
	// connect to postgres
	conn, err := postgres.NewClient(ctx, conf)
	if err != nil {
		log.Fatal().Str("service", "postgres").Err(err).Send()
	}
	log.Info().Msg("Postgress Conected!")

	// Migratons
	err = conn.Migrate()
	if err != nil {
		log.Fatal().Str("service", "postgres").Err(err).Send()
	}
	log.Info().Msg("Migrations succsess!")


}
