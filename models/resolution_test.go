package models_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"hungknow.com/blockchain/models"
)

func TestResolution(t *testing.T) {
	m1 := models.ResolutionM1
	m5 := models.ResolutionM5

	t.Run("BoundSeconds", func(t *testing.T) {
		cases := []struct {
			FromTime   int64
			ExpectedM1 []int64
			ExpectedM5 []int64
		}{
			{FromTime: 0, ExpectedM1: []int64{0, 60}, ExpectedM5: []int64{0, 300}},
			{FromTime: 1, ExpectedM1: []int64{0, 60}, ExpectedM5: []int64{0, 300}},
			{FromTime: 59, ExpectedM1: []int64{0, 60}, ExpectedM5: []int64{0, 300}},
			{FromTime: 60, ExpectedM1: []int64{60, 120}, ExpectedM5: []int64{0, 300}},
			{FromTime: 61, ExpectedM1: []int64{60, 120}, ExpectedM5: []int64{0, 300}},
		}
		for _, c := range cases {
			m1LowerBound, m2UpperBound := m1.BoundSeconds(c.FromTime)
			require.Equalf(t, c.ExpectedM1[0], m1LowerBound, "Expected %d, got %d", c.ExpectedM1[0], m1LowerBound)
			require.Equalf(t, c.ExpectedM1[1], m2UpperBound, "Expected %d, got %d", c.ExpectedM1[1], m2UpperBound)

			m5LowerBound, m5UpperBound := m5.BoundSeconds(c.FromTime)
			require.Equalf(t, c.ExpectedM5[0], m5LowerBound, "Expected %d, got %d", c.ExpectedM5[0], m5LowerBound)
			require.Equalf(t, c.ExpectedM5[1], m5UpperBound, "Expected %d, got %d", c.ExpectedM5[1], m5UpperBound)
		}
	})
}
