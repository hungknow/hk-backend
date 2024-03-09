package models

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
	CountBack     int `json:"countBack"`
	FromTimestamp int `json:"fromTimestamp"`
	ToTimestamp   int `json:"toTimestamp"`
}

type GetBarsResult struct {
	Bars     []Candle           `json:"bars"`
	Metadata HistoryMetadata `json:"meta"`
}