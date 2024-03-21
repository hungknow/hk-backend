package candleconv

import (
	"math"
	"time"

	"hungknow.com/blockchain/errors"
	"hungknow.com/blockchain/models"
)

func ConvertResolution(
	bars *models.Candles,
	toResolution models.Resolution,
) (*models.Candles, error) {
	if len(bars.Times) == 0 {
		return &models.Candles{}, nil
	}

	detectedResolution := bars.DetectResolution()
	if detectedResolution == models.ResolutionUnknown {
		return nil, errors.NewAppErrorf(errors.AppErrorInvalidParams, "Failed to detect resolution of the input bars")
	}

	if detectedResolution == toResolution {
		return bars, nil
	}

	if detectedResolution > toResolution {
		return nil, errors.NewAppErrorf(errors.AppErrorInvalidParams, "Failed to convert resolution from %s to %s", detectedResolution.String(), toResolution.String())
	}

	if detectedResolution.Seconds() > toResolution.Seconds() {
		return nil, errors.NewAppErrorf(errors.AppErrorInvalidParams, "The resolution %s of the input bars is less than as the new resolution %s", detectedResolution.String(), toResolution.String())
	}

	toResolutionBars := models.NewCandles()
	currentToResolutionBarsTimeIndex := -1
	previousTimeUpperBound := int64(-1)

	for index, timeOfCurrentBar := range bars.Times {
		timeOfCurrentBarLower, timeOfCurrentBarUpper := toResolution.BoundSeconds(timeOfCurrentBar.Unix())
		if timeOfCurrentBar.Unix() >= previousTimeUpperBound || currentToResolutionBarsTimeIndex < 0 {
			toResolutionBars = toResolutionBars.Push(
				time.Unix(timeOfCurrentBarLower, 0),
				bars.Opens[index],
				bars.Highs[index],
				bars.Lows[index],
				bars.Closes[index],
				bars.Vols[index],
			)
			previousTimeUpperBound = timeOfCurrentBarUpper
			currentToResolutionBarsTimeIndex = len(toResolutionBars.Times) - 1
		} else {
			toResolutionBars.Highs[currentToResolutionBarsTimeIndex] = math.Max(bars.Highs[index], toResolutionBars.Highs[currentToResolutionBarsTimeIndex])
			toResolutionBars.Lows[currentToResolutionBarsTimeIndex] = math.Min(bars.Lows[index], toResolutionBars.Lows[currentToResolutionBarsTimeIndex])
			toResolutionBars.Vols[currentToResolutionBarsTimeIndex] += bars.Vols[index]
			toResolutionBars.Closes[currentToResolutionBarsTimeIndex] = bars.Closes[index]
		}
	}

	return toResolutionBars, nil
}
