package models

type TimeRange struct {
	FromTimestamp int64 `json:"fromTimestamp"`
	ToTimestamp   int64 `json:"toTimestamp"`
	ToIsInclusive bool  `json:"toIsInclusive"`
}

func NewTimeRange(fromTimestamp int64, toTimestamp int64, toIsInclusive bool) TimeRange {
	return TimeRange{
		FromTimestamp: fromTimestamp,
		ToTimestamp:   toTimestamp,
		ToIsInclusive: toIsInclusive,
	}
}

func NewExclusiveTimeRange(fromTimestamp int64, toTimestamp int64) TimeRange {
	return TimeRange{
		FromTimestamp: fromTimestamp,
		ToTimestamp:   toTimestamp,
		ToIsInclusive: false,
	}
}

func NewInclusiveTimeRange(fromTimestamp int64, toTimestamp int64) TimeRange {
	return TimeRange{
		FromTimestamp: fromTimestamp,
		ToTimestamp:   toTimestamp,
		ToIsInclusive: true,
	}
}
