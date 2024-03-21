package candleconv_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"hungknow.com/blockchain/candleconv"
	"hungknow.com/blockchain/models"
)

func TestConvertResolution(t *testing.T) {
	bars := models.NewCandles().
		Push(time.Unix(1, 0), 1, 2, 3, 4, 6).
		Push(time.Unix(1+60, 0), 7, 8, 9, 10, 12).
		Push(time.Unix(1+60*2, 0), 13, 14, 15, 16, 18).
		Push(time.Unix(1+60*14, 0), 20, 21, 22, 23, 1).
		Push(time.Unix(1+60*15, 0), 24, 25, 26, 27, 1)

	expectedBarsM5 := models.NewCandles().
		Push(time.Unix(0, 0), 1, 14, 3, 16, 36).
		Push(time.Unix(60*10, 0), 20, 21, 22, 23, 1).
		Push(time.Unix(60*15, 0), 24, 25, 26, 27, 1)

	expectedBarsM15 := models.NewCandles().
		Push(time.Unix(0, 0), 1, 21, 3, 23, 37).
		Push(time.Unix(60*15, 0), 24, 25, 26, 27, 1)

	convertedBarsM5, err := candleconv.ConvertResolution(bars, models.ResolutionM5)
	require.NoError(t, err)
	require.Equal(t, expectedBarsM5, convertedBarsM5)

	convertedBarsM15, err := candleconv.ConvertResolution(bars, models.ResolutionM15)
	require.NoError(t, err)
	require.Equal(t, expectedBarsM15, convertedBarsM15)
}
