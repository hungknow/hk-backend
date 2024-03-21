package datafeed

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"

	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

type csvCandle struct {
	Ticker string
	Date   string
	Time   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type CandleCSVLoader struct {
	CsvReader *csv.Reader
}

func NewCandleCSVLoader(r io.Reader) *CandleCSVLoader {
	return &CandleCSVLoader{
		CsvReader: csv.NewReader(r),
	}
}

// func (o *CandleCSVLoader) GetMetadata() (*models.CandleMetadata, error) {
// 	return nil, nil
// }

func (o *CandleCSVLoader) rowToCandle(records []string) (*models.Candle, *errors.AppError) {
	date := records[1]
	hourminute := records[2]

	open, err := strconv.ParseFloat(records[3], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	high, err := strconv.ParseFloat(records[4], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	low, err := strconv.ParseFloat(records[5], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	close, err := strconv.ParseFloat(records[6], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	vol, err := strconv.ParseFloat(records[7], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}

	openTime, err := time.Parse("20060102 150405", date+" "+hourminute)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing time: %v", err)
	}

	return &models.Candle{
		Open:  open,
		High:  high,
		Low:   low,
		Close: close,
		Time:  openTime.Unix(),
		Vol:   vol,
	}, nil
}

func (o *CandleCSVLoader) GetCandles(fromTime time.Time, toTime time.Time) (*models.Candles, *errors.AppError) {
	candles := models.NewCandles()

	fromTimeIsZero := fromTime.IsZero()
	fromTimeUnix := fromTime.Unix()
	toTimeIsZero := toTime.IsZero()
	toTimeUnix := toTime.Unix()

	for {
		records, err := o.CsvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error reading csv: %v", err)
		}

		// Ignore header
		if records[0] == "ticker" {
			continue
		}

		candle, appErr := o.rowToCandle(records)
		if appErr != nil {
			return nil, appErr
		}

		if !toTimeIsZero && candle.Time > toTimeUnix {
			break
		}
		if fromTimeIsZero || candle.Time >= fromTimeUnix {
			candles.PushCandle(candle)
		}
	}

	return candles, nil
}
