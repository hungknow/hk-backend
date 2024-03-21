package main

import (
	"context"
	"os"
	"path"
	"time"

	"github.com/phuslu/log"
	"hungknow.com/blockchain/config"
	"hungknow.com/blockchain/datafeed"
	"hungknow.com/blockchain/db/dbstoresql"
	"hungknow.com/blockchain/logutils"
	"hungknow.com/blockchain/models"
)

// Read CSV files and insert candles into DB
func main() {
	logutils.SetupPhusluLog()

	ctx := context.Background()

	// Read DB config file
	appConfig, err := config.GetConfig()
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}

	// Create DB store
	sqlDB, err := dbstoresql.NewDBSQLStore(appConfig.DB)
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}

	// Load CSV files
	csvFilePath := path.Clean("data/candles_csv/XAUUSD_2023_01.csv")
	f, nerr := os.Open(csvFilePath)
	if nerr != nil {
		log.Panic().Msgf("%+v", err)
	}
	defer f.Close()

	candleCSVLoader := datafeed.NewCandleCSVLoader(f, datafeed.CsvHeaderTypeDotDateTime)
	candles, err := candleCSVLoader.GetCandles(time.Time{}, time.Time{})
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}

	candleLen := len(candles.Times)
	log.Debug().Msgf("Loaded %d candles", candleLen)

	if candleLen == 0 {
		log.Debug().Msgf("No candles to load")
		return
	}

	// Insert candles into DB
	forexCandleDB := sqlDB.ForexCandles()
	err = forexCandleDB.UpsertCandles(ctx, 1, models.ResolutionM1, candles)
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}

	log.Info().Msg("DONE!!!")
}
