package models

type SymbolSticker string

type SymbolInfo struct {
	ID int64 `json:"-"`
	SupportedResolutions []Resolution `json:"supportedResolutions"`
	// Unique symbol id It's an unique identifier for this particular symbol in your symbology.
	Sticker    SymbolSticker `json:"sticker"`
	SymbolType SymbolType    `json:"symbolType"`
}
