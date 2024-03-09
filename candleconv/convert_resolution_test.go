package candleconv_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"hungknow.com/blockchain/candleconv"
	"hungknow.com/blockchain/models"
)

func TestConvertResolution(t *testing.T) {
	bars := models.NewCandles().
		Push(1, 1, 2, 3, 4, 6).
		Push(1+60, 7, 8, 9, 10, 12).
		Push(1+60*2, 13, 14, 15, 16, 18)
	expectedBars := models.NewCandles().
		Push(0, 1, 14, 3, 16, 36)

	convertedBars, err := candleconv.ConvertResolution(bars, models.Resolution5M)
	require.NoError(t, err)
	require.Equal(t, expectedBars, convertedBars)
}
