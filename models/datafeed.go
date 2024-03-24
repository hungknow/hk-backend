package models

import (
	"time"

	"hungknow.com/blockchain/errors"
)

/*
	{
	    description: 'Apple Inc.',
	    exchange: 'NasdaqNM',
	    full_name: 'NasdaqNM:AAPL',
	    symbol: 'AAPL',
	    ticker: 'AAPL',
	    type: 'stock',
	}
*/
type SearchSymbolResultItem struct {
	Description string `json:"description"`
	Exchange    string `json:"exchange"`
	FullName    string `json:"full_name"`
	Symbol      string `json:"symbol"`
	Ticker      string `json:"ticker"`
	Type        string `json:"type"`
}

type PeriodParams struct {
	// the number of ohlc bar to return
	Count         int        `json:"countBack"`
	FromTimestamp time.Time  `json:"fromTimestamp"`
	ToTimestamp   time.Time  `json:"toTimestamp"`
	Resolution    Resolution `json:"resolution"`
}

func (p PeriodParams) IsValid() *errors.AppError {
	isZero := p.FromTimestamp.IsZero() || p.ToTimestamp.IsZero()
	if isZero {
		return errors.NewAppErrorf(errors.AppErrorInvalidParams, "FromTimestamp and ToTimestamp are required")
	}

	if p.FromTimestamp.Compare(p.ToTimestamp) >= 0 {
		return errors.NewAppErrorf(errors.AppErrorInvalidParams, "FromTimestamp must be greater than ToTimestamp")
	}

	return nil
}

func PeriodParamsPrepare(p *PeriodParams, resolution Resolution, maxBarCount int64) *errors.AppError {
	if p.FromTimestamp.IsZero() && p.ToTimestamp.IsZero() {
		return errors.NewAppErrorf(errors.AppErrorInvalidParams, "FromTimestamp and ToTimestamp are required")
	}
	maxPeriod := resolution.Seconds() * maxBarCount

	// The FromTimestamp is not filled, the ToTimestamp is filled,
	if p.FromTimestamp.IsZero() && !p.ToTimestamp.IsZero() {
		p.FromTimestamp = p.ToTimestamp.Add(time.Duration(-maxPeriod) * time.Second)
	}

	// The FromTimestamp is filled, the ToTimestamp isn't filled
	if !p.FromTimestamp.IsZero() && p.ToTimestamp.IsZero() {
		p.ToTimestamp = p.FromTimestamp.Add(time.Duration(maxPeriod) * time.Second)
	}

	return nil
}

type GetBarsResult struct {
	Candles  *Candles        `json:"candles"`
	Metadata HistoryMetadata `json:"meta"`
}
