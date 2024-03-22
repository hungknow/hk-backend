package hktrading_server

import (
	"context"
	"net/http"
	"time"

	"hungknow.com/blockchain/datafeed"
	"hungknow.com/blockchain/models"
	"hungknow.com/blockchain/symbols"
)

type APIService struct {
	symbolService symbols.SymbolService
	forexDataFeed *datafeed.ForexDataFeed
}

func NewAPIService(
	symbolService symbols.SymbolService,
	forexDataFeed *datafeed.ForexDataFeed,
) *APIService {
	return &APIService{
		symbolService: symbolService,
		forexDataFeed: forexDataFeed,
	}
}

func (s *APIService) GetCandles(ctx context.Context, symbol SymbolTicker, resolution Resolution, from int32, to int32) (ImplResponse, error) {
	symbolInfo, appErr := s.symbolService.GetSymbolInfoBySticker(string(symbol))
	if appErr != nil {
		return Response(http.StatusNotFound, HkError{Message: appErr.Error()}), nil
	}

	param := &models.PeriodParams{
		FromTimestamp: time.Unix(int64(from), 0),
		ToTimestamp:   time.Unix(int64(to), 0),
	}

	barResult, appErr := s.forexDataFeed.GetBars(ctx, symbolInfo, resolution.ToModel(), param)
	if appErr != nil {
		return Response(http.StatusNotFound, HkError{Message: appErr.Error()}), nil
	}

	return Response(http.StatusOK, barResult.Candles), nil
}
