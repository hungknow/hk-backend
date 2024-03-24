package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"hungknow.com/blockchain/models"
)

func TestPeriodParamsPrepare(t *testing.T) {
	periodParams := &models.PeriodParams{}
	targetResolution := models.ResolutionM1
	var maxBarCount int64 = 61

	err := models.PeriodParamsPrepare(periodParams, targetResolution, maxBarCount)
	require.ErrorContains(t, err, "FromTimestamp and ToTimestamp are required")
	
	fromTimestamp, serr :=  time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	require.NoError(t, serr)
	expectedToTimestamp, serr := time.Parse(time.RFC3339, "2021-01-01T01:01:00Z")
	require.NoError(t, serr)

	periodParams = &models.PeriodParams{
		FromTimestamp: fromTimestamp,
	}
	err = models.PeriodParamsPrepare(periodParams, targetResolution, maxBarCount)
	require.Nil(t, err)
	require.Equal(t, expectedToTimestamp, periodParams.ToTimestamp)

	periodParams = &models.PeriodParams{
		ToTimestamp: expectedToTimestamp,
	}
	err = models.PeriodParamsPrepare(periodParams, targetResolution, maxBarCount)
	require.Nil(t, err)
	require.Equal(t, fromTimestamp, periodParams.FromTimestamp)	
}