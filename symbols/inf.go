package symbols

import (
	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

const (
	SymbolNameXAUUSD models.SymbolSticker = "Mock:XAUUSD"
)

type SymbolService interface {
	GetSymbolInfoBySticker(symbolSticker string) (*models.SymbolInfo, *errors.AppError)
}
