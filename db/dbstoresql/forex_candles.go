package dbstoresql

import (
	"context"
	"strconv"
	"time"

	"github.com/phuslu/log"
	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

const (
	TradingForexCandlesTable = "trading.forex_candles"
)

func (o *DBSQLStore) UpsertCandles(ctx context.Context, symbol_id int, resolution models.Resolution, bars *models.Candles) *errors.AppError {
	if bars == nil {
		return errors.NewAppErrorf(errors.AppErrorInvalidParams, "candles is empty")
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	numFields := 8
	sqlStr := `INSERT INTO ` + TradingForexCandlesTable + `("symbol_id", "resolution", "open_time", "open", "close", "high", "low", "volume") VALUES `
	vals := make([]interface{}, 0)
	for timeIdx, openTime := range bars.Times {
		sqlStr += "("
		for j := 0; j < numFields; j++ {
			sqlStr += `$` + strconv.Itoa(timeIdx*numFields+j+1) + `,`
		}
		sqlStr = sqlStr[:len(sqlStr)-1] + `),`

		vals = append(vals, symbol_id, resolution.String(), time.Unix(openTime, 0).UTC(),
			bars.Opens[timeIdx], bars.Closes[timeIdx], bars.Highs[timeIdx], bars.Lows[timeIdx], bars.Vols[timeIdx])
	}

	// Remove the trailing comma
	sqlStr = sqlStr[:len(sqlStr)-1]

	sqlStr += ` ON CONFLICT ("symbol_id", "resolution", "open_time") DO UPDATE SET "open" = EXCLUDED.open, "close" = EXCLUDED.close, "high" = EXCLUDED.high, "low" = EXCLUDED.low, "volume" = EXCLUDED.volume`
	sqlStr += ";"

	log.Debug().Msgf("UpsertCandles Stmt: %s", sqlStr)
	stmt, err := tx.PrepareContext(ctx, sqlStr)
	if err != nil {
		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, vals...)
	if err != nil {
		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}
	if res == nil {
	}

	if err := tx.Commit(); err != nil {
		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}

	return nil
}

func (o *DBSQLStore) QueryCandles(ctx context.Context, symbol models.SymbolName, resolution models.Resolution, from time.Time, exclusiveTo time.Time) (*models.Candles, *errors.AppError) {
	panic("implement me")
}
