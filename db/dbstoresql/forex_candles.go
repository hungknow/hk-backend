package dbstoresql

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/phuslu/log"
	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

const (
	TradingForexCandlesTable = "trading.forex_candles"
)

func (o *DBSQLStore) prepareUpsertCandles(thresholdRows int) string {
	numFields := 8
	builder := strings.Builder{}
	builder.WriteString(`INSERT INTO ` + TradingForexCandlesTable + `("symbol_id", "resolution", "open_time", "open", "close", "high", "low", "volume") VALUES `)

	for rowIndex := 0; rowIndex < thresholdRows; rowIndex++ {
		builder.WriteString("(")
		for j := 0; j < numFields; j++ {
			builder.WriteByte('$')
			builder.WriteString(strconv.Itoa(rowIndex*numFields + j + 1))
			if j < numFields-1 {
				builder.WriteByte(',')
			}
		}
		builder.WriteString(")")
		if rowIndex < thresholdRows-1 {
			builder.WriteByte(',')
		}
	}

	builder.WriteString(` ON CONFLICT ("symbol_id", "resolution", "open_time") DO UPDATE SET "open" = EXCLUDED.open, "close" = EXCLUDED.close, "high" = EXCLUDED.high, "low" = EXCLUDED.low, "volume" = EXCLUDED.volume;`)

	return builder.String()
}

func (o *DBSQLStore) UpsertCandles(ctx context.Context, symbol_id int64, resolution models.Resolution, bars *models.Candles) *errors.AppError {
	if bars == nil {
		return errors.NewAppErrorf(errors.AppErrorInvalidParams, "candles is empty")
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	thresholdRows := 1000
	maxLen := len(bars.Times)
	rowIdx := 0

	var stmt *sql.Stmt
	rowThreshold := -1

	for rowIdx < maxLen {
		newRowThreshold := min(thresholdRows, maxLen-rowIdx)
		if newRowThreshold != rowThreshold {
			// Prepare the upsert statement
			sqlStr := o.prepareUpsertCandles(newRowThreshold)
			stmt, err = tx.PrepareContext(ctx, sqlStr)
			defer stmt.Close()
			if err != nil {
				return errors.NewAppErrorf(errors.AppDatabaseError, "%v: %s", err, sqlStr)
			}
		}
		rowThreshold = newRowThreshold

		vals := make([]interface{}, 0)

		maxThreshold := rowIdx + rowThreshold
		for ; rowIdx < maxThreshold; rowIdx++ {
			vals = append(vals, symbol_id, resolution.String(),
				bars.Times[rowIdx],
				bars.Opens[rowIdx], bars.Closes[rowIdx],
				bars.Highs[rowIdx], bars.Lows[rowIdx], bars.Vols[rowIdx])
		}

		_, err = stmt.ExecContext(ctx, vals...)
		if err != nil {
			return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}

	log.Info().Msgf("Upserted %d candles", rowIdx)

	return nil
}

func (o *DBSQLStore) QueryCandles(ctx context.Context, symbol_id int64, resolution models.Resolution, from time.Time, exclusiveTo time.Time, limit int64) (*models.Candles, *errors.AppError) {
	log.Debug().Msgf("QueryCandles: symbol_id=%d, resolution=%s, from=%s, exclusiveTo=%s, limit=%d", symbol_id, resolution.String(), from.UTC().String(), exclusiveTo.UTC().String(), limit)
	queryStr := `SELECT "open_time", "open", "close", "high", "low", "volume" FROM ` + TradingForexCandlesTable + " WHERE symbol_id = $1 AND resolution = $2 AND open_time >= $3 AND open_time < $4 ORDER BY open_time ASC "
	if limit > 0 {
		queryStr += " LIMIT " + strconv.FormatInt(limit, 10)
	}
	queryStr += ";"

	args := []interface{}{
		symbol_id,
		resolution.String(),
		from,
		exclusiveTo,
	}

	rows, err := o.db.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}
	defer rows.Close()

	candles := models.NewCandles()

	for rows.Next() {
		candle := models.Candle{}
		err := rows.Scan(&candle.Time, &candle.Open, &candle.Close, &candle.High, &candle.Low, &candle.Vol)
		if err != nil {
			return nil, errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
		}
		candles.PushCandle(&candle)
	}

	return candles, nil
}
