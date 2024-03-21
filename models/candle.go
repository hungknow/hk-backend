package models

type Candle struct {
	Close float64 `json:"close"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Open  float64 `json:"open"`
	Time  int64   `json:"time"`
	Vol   float64 `json:"vol"`
}

type Candles struct {
	Closes []float64 `json:"closes"`
	Highs  []float64 `json:"highs"`
	Lows   []float64 `json:"lows"`
	Opens  []float64 `json:"opens"`
	Times  []int64   `json:"times"`
	Vols   []float64 `json:"vols"`
}

func NewCandles() *Candles {
	return &Candles{
		Closes: make([]float64, 0),
		Highs:  make([]float64, 0),
		Lows:   make([]float64, 0),
		Opens:  make([]float64, 0),
		Times:  make([]int64, 0),
		Vols:   make([]float64, 0),
	}
}

func (o *Candles) DetectResolution() Resolution {
	if len(o.Times) < 2 {
		return ResolutionUnknown
	}
	return ResolutionFromSeconds(o.Times[1] - o.Times[0])
}

func (o *Candles) PushCandle(candle *Candle) *Candles {
	o.Closes = append(o.Closes, candle.Close)
	o.Highs = append(o.Highs, candle.High)
	o.Lows = append(o.Lows, candle.Low)
	o.Opens = append(o.Opens, candle.Open)
	o.Times = append(o.Times, candle.Time)
	o.Vols = append(o.Vols, candle.Vol)
	return o
}

func (o *Candles) Push(time int64, open float64, high float64, low float64, close float64, vol float64) *Candles {
	o.Times = append(o.Times, time)
	o.Opens = append(o.Opens, open)
	o.Highs = append(o.Highs, high)
	o.Lows = append(o.Lows, low)
	o.Closes = append(o.Closes, close)
	o.Vols = append(o.Vols, vol)
	return o
}

func (o *Candles) Len() int {
	return len(o.Times)
}

func (o *Candles) IsEmpty() bool {
	return len(o.Times) == 0
}

func (o *Candles) Slice(fromIdx int, uptoIdx int) *Candles {
	return &Candles{
		Closes: o.Closes[fromIdx:uptoIdx],
		Highs:  o.Highs[fromIdx:uptoIdx],
		Lows:   o.Lows[fromIdx:uptoIdx],
		Opens:  o.Opens[fromIdx:uptoIdx],
		Times:  o.Times[fromIdx:uptoIdx],
		Vols:   o.Vols[fromIdx:uptoIdx],
	}
}
