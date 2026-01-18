package model

type Stock struct {
	Chart `json:"chart"`
}

type Chart struct {
	Result []Result `json:"result"`
}

type Result struct {
	Indicators `json:"indicators"`
}

type Indicators struct {
	Quote []Quote `json:"quote"`
}

type Quote struct {
	Volume []float64 `json:"volume"`
	Close  []float64 `json:"close"`
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	High   []float64 `json:"high"`
}
