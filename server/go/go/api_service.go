package hktrading_server

import (
	"context"
	"net/http"
)

type APIService struct {

}

func NewAPIService() *APIService {
	return &APIService{}
}

func (s *APIService) GetCandles(ctx context.Context, symbol SymbolTicker, resolution Resolution, from int32, to int32) (ImplResponse, error) {
	return Response(http.StatusOK, Candles{}), nil
}