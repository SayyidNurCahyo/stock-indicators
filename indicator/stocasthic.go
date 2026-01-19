package indicator

import (
	"njajal/model"

	"github.com/samber/lo"
)

func kLine(High, Low, Close []float64, period int) float64 {
	var highestHigh float64
	var lowestLow float64
	var close float64

	for i := (len(Close) - period - 15); i < (len(Close) - period); i++ {
		if highestHigh == 0 {
			highestHigh = High[i]
			lowestLow = Low[i]
			close = Close[i]
			continue
		}
		if High[i] > highestHigh {
			highestHigh = High[i]
		}
		if Low[i] < lowestLow {
			lowestLow = Low[i]
		}
		close = Close[i]
	}

	return ((close - lowestLow) / (highestHigh - lowestLow)) * 100
}

func Stochastic(data model.Quote, period int) (float64, float64) {
	high := lo.Filter(data.High, func(val float64, _ int) bool {
		return val != 0
	})
	low := lo.Filter(data.Low, func(val float64, _ int) bool {
		return val != 0
	})
	close := lo.Filter(data.Close, func(val float64, _ int) bool {
		return val != 0
	})
	if len(high) < 18 || len(low) < 18 || len(close) < 18 {
		return 0, 0
	}
	dLine := (kLine(high, low, close, 0+period) + kLine(high, low, close, 1+period) + kLine(high, low, close, 2+period)) / 3

	return kLine(high, low, close, 0+period), dLine
}
