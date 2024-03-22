package symbols

import (
	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

type SymbolManager struct {
}

func NewSymbolManager() *SymbolManager {
	return &SymbolManager{}
}

func (o *SymbolManager) GetSymbolInfoBySticker(symbolSticker string) (*models.SymbolInfo, *errors.AppError) {
	symbolInfo := &models.SymbolInfo{
		ID:                   1,
		SupportedResolutions: []models.Resolution{models.ResolutionM1},
		Sticker:              SymbolNameXAUUSD,
		SymbolType:           models.SymbolTypeForex,
	}
	return symbolInfo, nil
}
