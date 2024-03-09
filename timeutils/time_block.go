package timeutils

func CreateTimeBlock(
	fromSeconds int64,
	toSeconds int64,
	secondDuration int64,
) [][]int64 {
	// Create the time blocks
	res := make([][]int64, 0)

	fromTimeLower := fromSeconds - (fromSeconds % secondDuration)
	if fromTimeLower+secondDuration > toSeconds {
		return append(res, []int64{fromSeconds, toSeconds})
	}
	res = append(res, []int64{fromSeconds, fromTimeLower + secondDuration})
	fromTimeLower += secondDuration

	for fromTimeLower < toSeconds {
		// The to time is exclusive
		toTimeUpper := min(toSeconds, fromTimeLower+secondDuration)
		// The from time is inclusive
		res = append(res, []int64{fromTimeLower, toTimeUpper})
		fromTimeLower = toTimeUpper
	}

	return res
}
