package timeutils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"hungknow.com/blockchain/models"
	"hungknow.com/blockchain/timeutils"
)

func TestCreateTimeBlock(t *testing.T) {
	cases := []struct {
		params []int64
		Result []models.TimeRange
	}{
		{params: []int64{0, 10, 5}, Result: []models.TimeRange{models.NewExclusiveTimeRange(0, 5), models.NewExclusiveTimeRange(5, 10)}},
		{params: []int64{1, 9, 5}, Result: []models.TimeRange{models.NewExclusiveTimeRange(1, 5), models.NewExclusiveTimeRange(5, 9)}},
		{params: []int64{1, 9, 10}, Result: []models.TimeRange{models.NewExclusiveTimeRange(1, 9)}},
	}

	for _, c := range cases {
		res := timeutils.CreateTimeBlock(c.params[0], c.params[1], c.params[2])
		require.Equal(t, c.Result, res, "Expected %v, got %v", c.Result, res)
	}
}
