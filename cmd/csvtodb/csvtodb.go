package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"hungknow.com/blockchain/config"
	"hungknow.com/blockchain/db/dbstoresql"
	"hungknow.com/blockchain/models"
)

func main() {
	ctx := context.Background()

	// Read DB config file
	appConfig, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	// Create DB store
	sqlDB, err := dbstoresql.NewDBSQLStore(appConfig.DB)
	if err != nil {
		panic(err)
	}

	// Load CSV files
	forexCandleDB := sqlDB.ForexCandles()
	err = forexCandleDB.UpsertCandles(ctx, "XAUUSD", models.ResolutionM1, nil)
	if err != nil {
		panic(err)
	}

	logrus.Info("DONE!!!")
}
