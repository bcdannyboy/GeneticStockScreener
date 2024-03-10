package FMP

import "github.com/spacecodewor/fmpcloud-go/objects"

func (f *FMPAPI) GetPriceChange(Symbol string) ([]objects.StockPriceChange, error) {
	return f.APIClient.Stock.PriceChange(Symbol)
}
