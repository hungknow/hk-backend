package models

// HistoryMetadata is a struct that represents the metadata of a historical data request.
type HistoryMetadata struct {
	NextTime int  `json:"nextTime"`
	NoData   bool `json:"noData"`
}
