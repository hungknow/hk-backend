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

// func (o *DBSQLStore) SubmitCandles(ctx context.Context, stmt *sql.Stmt, symbol_id int, resolution models.Resolution, barIndex int, bars *models.Candles, thresholdRows int) *errors.AppError {
// 	vals := make([]interface{}, 0)

// 	sqlStr := ""
// 	numFields := 8
// 	maxLen := len(bars.Times)
// 	for rowIndex := 0; rowIndex < thresholdRows && barIndex < maxLen; rowIndex++ {
// 		sqlStr += "("
// 		for j := 0; j < numFields; j++ {
// 			sqlStr += `$` + strconv.Itoa(rowIndex*numFields+j+1) + `,`
// 		}
// 		sqlStr = sqlStr[:len(sqlStr)-1] + `),`
// 		barIndex++

// 		vals = append(vals, symbol_id, resolution.String(),
// 			time.Unix(bars.Times[barIndex], 0).UTC(),
// 			bars.Opens[barIndex], bars.Closes[barIndex],
// 			bars.Highs[barIndex], bars.Lows[barIndex], bars.Vols[barIndex])
// 	}

// 	log.Debug().Msgf("UpsertCandles Stmt: %s", sqlStr)

// 	res, err := stmt.ExecContext(ctx, vals...)
// 	if err != nil {
// 		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
// 	}

// 	return nil
// }

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
			if err != nil {
				return errors.NewAppErrorf(errors.AppDatabaseError, "%v: %s", err, sqlStr)
			}
		}
		rowThreshold = newRowThreshold

		vals := make([]interface{}, 0)

		maxThreshold := rowIdx + rowThreshold
		for ; rowIdx < maxThreshold; rowIdx++ {
			vals = append(vals, symbol_id, resolution.String(),
				time.Unix(bars.Times[rowIdx], 0).UTC(),
				bars.Opens[rowIdx], bars.Closes[rowIdx],
				bars.Highs[rowIdx], bars.Lows[rowIdx], bars.Vols[rowIdx])
		}

		_, err = stmt.ExecContext(ctx, vals...)
		if err != nil {
			return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
		}
	}

	// for timeIdx := 0; timeIdx < maxLen; {
	// 	sqlStr := `INSERT INTO ` + TradingForexCandlesTable + `("symbol_id", "resolution", "open_time", "open", "close", "high", "low", "volume") VALUES `

	// 	for rowIndex := 0; rowIndex < thresholdRows && timeIdx < maxLen; rowIndex++ {
	// 		sqlStr += "("
	// 		for j := 0; j < numFields; j++ {
	// 			sqlStr += `$` + strconv.Itoa(rowIndex*numFields+j+1) + `,`
	// 		}
	// 		sqlStr = sqlStr[:len(sqlStr)-1] + `),`
	// 		timeIdx++

	// 	}

	// 	// Remove the trailing comma
	// 	sqlStr = sqlStr[:len(sqlStr)-1]

	// 	sqlStr += ` ON CONFLICT ("symbol_id", "resolution", "open_time") DO UPDATE SET "open" = EXCLUDED.open, "close" = EXCLUDED.close, "high" = EXCLUDED.high, "low" = EXCLUDED.low, "volume" = EXCLUDED.volume`
	// 	sqlStr += ";"

	// 	log.Debug().Msgf("UpsertCandles Stmt: %s", sqlStr)
	// 	stmt, err := tx.PrepareContext(ctx, sqlStr)
	// 	if err != nil {
	// 		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	// 	}
	// 	defer stmt.Close()

	// 	if res == nil {
	// 	}
	// }

	if err := tx.Commit(); err != nil {
		return errors.NewAppErrorf(errors.AppDatabaseError, "%v", err)
	}

	log.Info().Msgf("Upserted %d candles", rowIdx)

	return nil
}

func (o *DBSQLStore) QueryCandles(ctx context.Context, symbol models.SymbolName, resolution models.Resolution, from time.Time, exclusiveTo time.Time) (*models.Candles, *errors.AppError) {
	panic("implement me")
}
