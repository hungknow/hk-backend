package datafeed

import (
	"context"
	"time"

	"hungknow.com/blockchain/db/dbstore"
	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
	"hungknow.com/blockchain/symbols"
)

type ForexDataFeed struct {
	dbStore       dbstore.DBStore
	symbolManager symbols.SymbolService
}

func NewForexDataFeed(dbStore dbstore.DBStore, symbolManager symbols.SymbolService) *ForexDataFeed {
	return &ForexDataFeed{
		dbStore:       dbStore,
		symbolManager: symbolManager,
	}
}

func (o *ForexDataFeed) GetBars(
	ctx context.Context,
	symbolInfo *models.SymbolInfo,
	resolution models.Resolution,
	periodParams *models.PeriodParams,
) (*models.GetBarsResult, *errors.AppError) {
	toTimestamp := periodParams.ToTimestamp
	if periodParams.ToTimestamp.IsZero() {
		toTimestamp = time.Now()
	}

	// symbolManager.GetSymbolInfoBySticker(symbolInfo.Sticker)
	candles, err := o.dbStore.ForexCandles().QueryCandles(ctx, symbolInfo.ID, resolution,
		periodParams.FromTimestamp, toTimestamp, 1500)
	if err != nil {
		return nil, err
	}

	return &models.GetBarsResult{
		Candles: candles,
		Metadata: models.HistoryMetadata{
			NextTime: 0,
			NoData:   true,
		},
	}, nil
}
