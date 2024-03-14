package FMP

import "github.com/spacecodewor/fmpcloud-go/objects"

func (fmp *FMPAPI) GetHistoricalPrices(ticker string) (*objects.StockDailyCandleList, error) {
	return fmp.APIClient.Stock.DailyChangeAndVolume(ticker)
}
