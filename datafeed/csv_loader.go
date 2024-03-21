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

type CsvHeaderType int

const (
	// 2023.01.03,01:00,1826.97,1827.18,1826.39,1826.43,71
	CsvHeaderTypeDotDateTime CsvHeaderType = iota
	// XAUUSD,20220608,000000,1849,1849.1,1849,1849,4
	CsvHeaderTypeTickerDateTime
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
	CsvReader  *csv.Reader
	headerType CsvHeaderType
}

func NewCandleCSVLoader(r io.Reader, headerType CsvHeaderType) *CandleCSVLoader {
	return &CandleCSVLoader{
		CsvReader:  csv.NewReader(r),
		headerType: headerType,
	}
}

func (o *CandleCSVLoader) CsvHeaderTypeDotDateTimeParseRow(records []string) (*models.Candle, *errors.AppError) {
	openTime, err := time.Parse("2006.01.02 15:04", records[0]+" "+records[1])
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing time: %v", err)
	}
	open, err := strconv.ParseFloat(records[2], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	high, err := strconv.ParseFloat(records[3], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	low, err := strconv.ParseFloat(records[4], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	close, err := strconv.ParseFloat(records[5], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	vol, err := strconv.ParseFloat(records[6], 64)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppLoadDataError, "error parsing data: %v, %s", err, strings.Join(records, ","))
	}
	return &models.Candle{
		Open:  open,
		High:  high,
		Low:   low,
		Close: close,
		Time:  openTime,
		Vol:   vol,
	}, nil
}

func (o *CandleCSVLoader) CsvHeaderTypeTickerDateTimeParseRow(records []string) (*models.Candle, *errors.AppError) {
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
		Time:  openTime,
		Vol:   vol,
	}, nil
}

func (o *CandleCSVLoader) GetCandles(fromTime time.Time, toTime time.Time) (*models.Candles, *errors.AppError) {
	candles := models.NewCandles()

	fromTimeIsZero := fromTime.IsZero()
	fromTimeUnix := fromTime
	toTimeIsZero := toTime.IsZero()
	toTimeUnix := toTime

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

		var candle *models.Candle
		var appErr *errors.AppError
		switch o.headerType {
		case CsvHeaderTypeDotDateTime:
			candle, appErr = o.CsvHeaderTypeDotDateTimeParseRow(records)
		case CsvHeaderTypeTickerDateTime:
			candle, appErr = o.CsvHeaderTypeTickerDateTimeParseRow(records)
		default:
			appErr = errors.NewAppErrorf(errors.AppLoadDataError, "unknown header type: %v", o.headerType)
		}
		if appErr != nil {
			return nil, appErr
		}

		if !toTimeIsZero && candle.Time.After(toTimeUnix) {
			break
		}
		if fromTimeIsZero || candle.Time.Compare(fromTimeUnix) >= 0 {
			candles.PushCandle(candle)
		}
	}

	return candles, nil
}
