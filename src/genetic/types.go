package genetic

import (
	"github.com/bcdannyboy/GeneticStockScreener/src/FMP"
	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

type Individual struct {
	Symbol           string
	Fundamentals     *FMP.CompanyValuationInfo
	PriceChangeScore float64
	Weight           *genetic_weight.Weight
	FundamentalScore float64
}
