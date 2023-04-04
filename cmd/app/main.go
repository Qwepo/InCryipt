package main

import (
	"context"
	"fmt"

	internal "github.com/Qwepo/InCryipt/Internal"
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
	log.Info().Msg("Start")
	_, err = postgres.NewClient(ctx, conf)
	if err != nil {
		fmt.Println(err)
	}

}
