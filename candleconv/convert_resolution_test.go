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
		Push(1+60*2, 13, 14, 15, 16, 18).
		Push(1+60*14, 20, 21, 22, 23, 1).
		Push(1+60*15, 24, 25, 26, 27, 1)

	expectedBarsM5 := models.NewCandles().
		Push(0, 1, 14, 3, 16, 36).
		Push(60*10, 20, 21, 22, 23, 1).
		Push(60*15, 24, 25, 26, 27, 1)

	expectedBarsM15 := models.NewCandles().
		Push(0, 1, 21, 3, 23, 37).
		Push(60*15, 24, 25, 26, 27, 1)

	convertedBarsM5, err := candleconv.ConvertResolution(bars, models.Resolution5M)
	require.NoError(t, err)
	require.Equal(t, expectedBarsM5, convertedBarsM5)

	convertedBarsM15, err := candleconv.ConvertResolution(bars, models.Resolution15M)
	require.NoError(t, err)
	require.Equal(t, expectedBarsM15, convertedBarsM15)
}
