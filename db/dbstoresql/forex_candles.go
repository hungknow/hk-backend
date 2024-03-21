package dbstoresql

import (
	"context"
	"time"

	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

func (o *DBSQLStore) UpsertCandles(ctx context.Context, symbol models.SymbolName, resolution models.Resolution, bars *models.Candles) *errors.AppError {
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.NewAppErrorf(errors.AppInternalError, "%v", err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	tx.ExecContext(ctx, "INSERT INTO a (id, x) VALUES() ON CONFLICT (id) DO UPDATE SET x = EXCLUDED.x")

	return nil
}

func (o *DBSQLStore) QueryCandles(ctx context.Context, symbol models.SymbolName, resolution models.Resolution, from time.Time, exclusiveTo time.Time) (*models.Candles, *errors.AppError) {
	panic("implement me")
}
