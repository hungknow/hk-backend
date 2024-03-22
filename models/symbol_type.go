package models

type SymbolType string

const (
	SymbolTypeForex  SymbolType = "forex"
	SymbolTypeCrypto SymbolType = "crypto"
	SymbolTypeStock  SymbolType = "stock"
)
