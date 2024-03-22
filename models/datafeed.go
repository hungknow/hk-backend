package models

import "time"

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
	CountBack     int       `json:"countBack"`
	FromTimestamp time.Time `json:"fromTimestamp"`
	ToTimestamp   time.Time `json:"toTimestamp"`
}

type GetBarsResult struct {
	Candles  *Candles        `json:"candles"`
	Metadata HistoryMetadata `json:"meta"`
}
