package FMP

import (
	"fmt"

	"github.com/bcdannyboy/GeneticStockScreener/src/utils"
	fmpcloud "github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

type FMPAPI struct {
	APIClient *fmpcloud.APIClient
}

func NewFMPAPI(apiKey string) (*FMPAPI, error) {
	APIClient, err := fmpcloud.NewAPIClient(fmpcloud.Config{APIKey: apiKey})
	if err != nil {
		return nil, fmt.Errorf("error creating FMP API client: %s", err)
	}

	return &FMPAPI{APIClient: APIClient}, nil
}

func (fmp *FMPAPI) GetValuationInfo(Symbol string, Period objects.CompanyValuationPeriod) (*CompanyValuationInfo, *objects.StockDailyCandleList, error) {
	KeyMetrics, KeyMetricsTTM, err := fmp.GetKeyMetrics(Symbol, Period)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting key metrics: %s", err)
	}

	AvgKeyMetrics := GetAverageKeyMetrics(KeyMetrics)

	Ratios, RatiosTTM, err := fmp.GetRatios(Symbol, Period)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting financial ratios: %s", err)
	}

	AvgRatios := GetAverageFinancialRatios(Ratios)

	CashFlowStatementGrowth, err := fmp.GetCashFlowStatementGrowth(Symbol, Period)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting cash flow growth: %s", err)
	}

	AvgCashFlowStatementGrowth := GetAverageCashFlowStatementGrowth(CashFlowStatementGrowth)

	IncomeStatementGrowth, err := fmp.GetIncomeStatementGrowth(Symbol, Period)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting income statement growth: %s", err)
	}

	AvgIncomeStatementGrowth := GetAverageIncomeStatementGrowth(IncomeStatementGrowth)

	BalanceSheetStatementGrowth, err := fmp.GetBalanceSheetStatementGrowth(Symbol, Period)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting balance sheet statement growth: %s", err)
	}

	AvgBalanceSheetStatementGrowth := GetAverageBalanceSheetStatementGrowth(BalanceSheetStatementGrowth)

	FinancialStatementsGrowth, err := fmp.GetFinancialStatementGrowth(Symbol, Period)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting financial statements growth: %s", err)
	}

	AvgFinancialStatementsGrowth := GetAverageFinancialStatementGrowth(FinancialStatementsGrowth)

	// only accept companies with full data

	if len(KeyMetrics) == 0 || len(Ratios) == 0 || len(CashFlowStatementGrowth) == 0 || len(IncomeStatementGrowth) == 0 || len(BalanceSheetStatementGrowth) == 0 || len(FinancialStatementsGrowth) == 0 {
		return nil, nil, fmt.Errorf("error getting valuation info: some data is missing (arrays)")
	}

	if utils.IsZeroValue(KeyMetricsTTM) || utils.IsZeroValue(AvgKeyMetrics) || utils.IsZeroValue(RatiosTTM) || utils.IsZeroValue(AvgRatios) || utils.IsZeroValue(AvgCashFlowStatementGrowth) || utils.IsZeroValue(AvgIncomeStatementGrowth) || utils.IsZeroValue(AvgBalanceSheetStatementGrowth) || utils.IsZeroValue(AvgFinancialStatementsGrowth) {
		return nil, nil, fmt.Errorf("error getting valuation info: some data is missing (structs)")
	}

	priceList, err := fmp.GetHistoricalPrices(Symbol)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting historical prices: %s", err)
	}

	return &CompanyValuationInfo{
		KeyMetrics:                     KeyMetrics,
		KeyMetricsTTM:                  KeyMetricsTTM,
		AvgKeyMetrics:                  AvgKeyMetrics,
		Ratios:                         Ratios,
		RatiosTTM:                      RatiosTTM,
		AvgRatios:                      AvgRatios,
		CashFlowStatementGrowth:        CashFlowStatementGrowth,
		AvgCashFlowStatementGrowth:     AvgCashFlowStatementGrowth,
		IncomeStatementGrowth:          IncomeStatementGrowth,
		AvgIncomeStatementGrowth:       AvgIncomeStatementGrowth,
		BalanceSheetStatementGrowth:    BalanceSheetStatementGrowth,
		AvgBalanceSheetStatementGrowth: AvgBalanceSheetStatementGrowth,
		FinancialStatementsGrowth:      FinancialStatementsGrowth,
		AvgFinancialStatementsGrowth:   AvgFinancialStatementsGrowth,
	}, priceList, nil
}
