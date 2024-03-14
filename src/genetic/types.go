package genetic

import (
	"github.com/bcdannyboy/GeneticStockScreener/src/FMP"
	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

type Individual struct {
	Symbol           string
	Fundamentals     *FMP.CompanyValuationInfo
	Weight           *genetic_weight.Weight
	FundamentalScore float64
	StockCandles     *objects.StockDailyCandleList
}
