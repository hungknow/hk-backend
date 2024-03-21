package models

type Resolution int

const (
	ResolutionUnknown Resolution = 0
	ResolutionS1      Resolution = 1
	ResolutionM1      Resolution = 60
	ResolutionM5      Resolution = 300
	ResolutionM15     Resolution = 900
	ResolutionH1      Resolution = 3600
	ResolutionH4      Resolution = 14400
	ResolutionD1      Resolution = 86400
	ResolutionW1      Resolution = 604800
)

func ResolutionFromSeconds(seconds int64) Resolution {
	switch seconds {
	case 1:
		return ResolutionS1
	case 60:
		return ResolutionM1
	case 300:
		return ResolutionM5
	case 900:
		return ResolutionM15
	case 3600:
		return ResolutionH1
	case 14400:
		return ResolutionH4
	case 86400:
		return ResolutionD1
	case 604800:
		return ResolutionW1
	default:
		return ResolutionUnknown
	}
}

func (r Resolution) String() string {
	switch r {
	case ResolutionS1:
		return "S1"
	case ResolutionM1:
		return "M1"
	case ResolutionM5:
		return "M5"
	case ResolutionM15:
		return "M15"
	case ResolutionH1:
		return "H1"
	case ResolutionH4:
		return "H4"
	case ResolutionD1:
		return "D1"
	case ResolutionW1:
		return "W1"
	default:
		return "Unknown"
	}
}

func (r Resolution) Seconds() int64 {
	return int64(r)
}

// Return Range [fromTime, toTime)
func (r Resolution) BoundSeconds(time int64) (int64, int64) {
	secondDurations := r.Seconds()
	lowerBound := time - time%secondDurations
	return lowerBound, lowerBound + secondDurations
}
