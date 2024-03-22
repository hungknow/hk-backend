package datafeed

import (
	"context"

	"hungknow.com/blockchain/models"
)

type Datafeed interface {
	GetBars(
		ctx context.Context,
		symbolInfo *models.SymbolInfo,
		resolution models.Resolution,
		periodParams *models.PeriodParams,
	) (*models.GetBarsResult, error)
	SearchSymbols(
		ctx context.Context,
		userInput string, // Text entered by user in the symbol search field
		exchange string, // The requested exchange. Empty value means no filter was specified
		symbolType string, // Type of symbol. Empty value means no filter was specified
	) []*models.SearchSymbolResultItem
	ResolveSymbol(
		ctx context.Context,
		symbolName string,
	) (*models.SymbolInfo, error)
}
