package models

type Exchange struct {
	Desc  string `json:"desc"`  // Description of the exchange
	Name  string `json:"name"`  // Name of the exchange
	Value string `json:"value"` // Value to be passed as the exchange argument to searchSymbols
}
