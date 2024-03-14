package FMP

import "github.com/spacecodewor/fmpcloud-go/objects"

//	type PriceCharts struct {
//		1D
//	}
type ValuationScore struct {
	AltmanZScore   float64 `json:"AltmanZScore"`
	PiotroskiScore float64 `json:"PiotroskiScore"`
}
type CompanyValuationInfo struct {
	KeyMetrics                     []objects.KeyMetrics
	AvgKeyMetrics                  objects.KeyMetrics
	KeyMetricsTTM                  []objects.KeyMetricsTTM
	Ratios                         []objects.FinancialRatios
	AvgRatios                      objects.FinancialRatios
	RatiosTTM                      []objects.FinancialRatiosTTM
	CashFlowStatementGrowth        []objects.CashFlowStatementGrowth
	AvgCashFlowStatementGrowth     objects.CashFlowStatementGrowth
	IncomeStatementGrowth          []objects.IncomeStatementGrowth
	AvgIncomeStatementGrowth       objects.IncomeStatementGrowth
	BalanceSheetStatementGrowth    []objects.BalanceSheetStatementGrowth
	AvgBalanceSheetStatementGrowth objects.BalanceSheetStatementGrowth
	FinancialStatementsGrowth      []objects.FinancialStatementsGrowth
	AvgFinancialStatementsGrowth   objects.FinancialStatementsGrowth
}
