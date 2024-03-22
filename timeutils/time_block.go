package timeutils

import "hungknow.com/blockchain/models"

func CreateTimeBlock(
	fromSeconds int64,
	toSeconds int64,
	secondDuration int64,
) []models.TimeRange {
	// Create the time blocks
	res := make([]models.TimeRange, 0)

	fromTimeLower := fromSeconds - (fromSeconds % secondDuration)
	if fromTimeLower+secondDuration > toSeconds {
		return append(res, models.NewExclusiveTimeRange(fromSeconds, toSeconds))
	}
	res = append(res, models.NewExclusiveTimeRange(fromSeconds, fromTimeLower+secondDuration))
	fromTimeLower += secondDuration

	for fromTimeLower < toSeconds {
		// The to time is exclusive
		toTimeUpper := min(toSeconds, fromTimeLower+secondDuration)
		// The from time is inclusive
		res = append(res, models.NewExclusiveTimeRange(fromTimeLower, toTimeUpper))
		fromTimeLower = toTimeUpper
	}

	return res
}
