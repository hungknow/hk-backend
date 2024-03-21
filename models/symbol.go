package models

type SymbolName string

const (
	SymbolNameXAUUSD SymbolName = "XAUUSD"
)

type SymbolInfo struct {
	BaseNames            []string     `json:"baseNames"`
	SupportedResolutions []Resolution `json:"supportedResolutions"`
	// Unique symbol id It's an unique identifier for this particular symbol in your symbology.
	Ticket     string     `json:"ticket"`
	SymbolType SymbolType `json:"symbolType"`
}
