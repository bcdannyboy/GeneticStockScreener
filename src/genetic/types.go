package genetic

import (
	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

type Individual struct {
	Weight         *genetic_weight.Weight
	PortfolioScore float64
}

type TickerScore struct {
	Ticker string
	Score  float64
}
