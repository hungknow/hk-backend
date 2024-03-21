package dbstore

import (
	"context"
	"time"

	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

type ForexCandles interface {
	UpsertCandles(ctx context.Context, symbol models.SymbolName, resolution models.Resolution, bars *models.Candles) *errors.AppError
	QueryCandles(ctx context.Context, symbol models.SymbolName, resolution models.Resolution, from time.Time, exclusiveTo time.Time) (*models.Candles, *errors.AppError)
}

type DBStore interface {
	ForexCandles() ForexCandles
}
