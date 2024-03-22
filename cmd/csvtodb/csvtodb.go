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
	appConfig, err := config.GetConfig("../../")
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}

	// Create DB store
	sqlDB, err := dbstoresql.NewDBSQLStore(appConfig.DB)
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}

	// Load CSV files
	csvFilePath := path.Clean("data/candles_csv/XAUUSD/XAUUSD_2023_01.csv")
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

	startingTime := candles.GetTimeAt(0)
	endingTime := candles.GetTimeAt(candleLen - 1)

	// Get the Forex candles store
	forexCandleDB := sqlDB.ForexCandles()

	// Insert candles into DB
	// err = forexCandleDB.UpsertCandles(ctx, 1, models.ResolutionM1, candles)
	// if err != nil {
	// 	log.Panic().Msgf("%+v", err)
	// }

	queriedCandles, err := forexCandleDB.QueryCandles(ctx, 1, models.ResolutionM1, startingTime, endingTime.Add(time.Second), -1)
	if err != nil {
		log.Panic().Msgf("%+v", err)
	}
	if queriedCandles.Len() != candleLen {
		log.Panic().Msgf("queriedCandles.Len() %d != candleLen %d", queriedCandles.Len(), candleLen)
	}
	if queriedCandles.Times[0] != candles.Times[0] {
		log.Panic().Msgf("queriedCandles.Times[0] %s != candles.Times[0] %s", queriedCandles.Times[0].String(), candles.Times[0].String())
	}

	log.Info().Msg("DONE!!!")
}
