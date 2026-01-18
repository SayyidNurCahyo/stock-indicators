package indicator

import "njajal/model"

func kLine(data model.Quote, period int) float64 {
	var highestHigh float64
	var lowestLow float64
	var close float64

	for i := 0; i < (len(data.Volume) - period); i++ {
		if data.High[i] == 0 || data.Low[i] == 0 || data.Close[i] == 0 {
			continue
		}
		if highestHigh == 0 {
			highestHigh = data.High[i]
			lowestLow = data.Low[i]
			close = data.Close[i]
			continue
		}
		if data.High[i] > highestHigh {
			highestHigh = data.High[i]
		}
		if data.Low[i] < lowestLow {
			lowestLow = data.Low[i]
		}
		close = data.Close[i]
	}

	return ((close - lowestLow) / (highestHigh - lowestLow)) * 100
}

func Stochastic(data model.Quote) (float64, float64) {
	dLine := (kLine(data, 0) + kLine(data, 1) + kLine(data, 2)) / 3

	return kLine(data, 0), dLine
}