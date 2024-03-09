package models

type Resolution int

const (
	ResolutionUnknown Resolution = 0
	Resolution1S      Resolution = 1
	Resolution1M      Resolution = 60
	Resolution5M      Resolution = 300
	Resolution15M     Resolution = 900
	Resolution1H      Resolution = 3600
	Resolution4H      Resolution = 14400
	Resolution1D      Resolution = 86400
	Resolution1W      Resolution = 604800
)

func ResolutionFromSeconds(seconds int64) Resolution {
	switch seconds {
	case 1:
		return Resolution1S
	case 60:
		return Resolution1M
	case 300:
		return Resolution5M
	case 900:
		return Resolution15M
	case 3600:
		return Resolution1H
	case 14400:
		return Resolution4H
	case 86400:
		return Resolution1D
	case 604800:
		return Resolution1W
	default:
		return ResolutionUnknown
	}
}

func (r Resolution) String() string {
	switch r {
	case Resolution1S:
		return "1S"
	case Resolution1M:
		return "1M"
	case Resolution5M:
		return "5M"
	case Resolution15M:
		return "15M"
	case Resolution1H:
		return "1H"
	case Resolution4H:
		return "4H"
	case Resolution1D:
		return "1D"
	case Resolution1W:
		return "1W"
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
