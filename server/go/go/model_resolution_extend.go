package hktrading_server

import "hungknow.com/blockchain/models"

func (v Resolution) ToModel() models.Resolution {
	switch v {
	case M1:
		return models.ResolutionM1
	case M5:
		return models.ResolutionM5
	case M15:
		return models.ResolutionM15
	case M30:
		return models.ResolutionM30
	case H1:
		return models.ResolutionH1
	case H4:
		return models.ResolutionH4
	case D1:
		return models.ResolutionD1
	case W1:
		return models.ResolutionW1
	default:
		return models.ResolutionUnknown
	}
}
