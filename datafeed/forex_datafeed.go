package datafeed

import (
	"context"

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
	var maxCount int64 = 200
	err := models.PeriodParamsPrepare(periodParams, resolution, maxCount)
	if err != nil {
		return nil, err
	}

	candles, err := o.dbStore.ForexCandles().QueryCandles(ctx, symbolInfo.ID, resolution,
		periodParams.FromTimestamp, periodParams.ToTimestamp, maxCount)
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
