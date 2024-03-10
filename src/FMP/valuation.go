package FMP

import (
	"fmt"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

func (fmp *FMPAPI) GetKeyMetrics(Symbol string, Period objects.CompanyValuationPeriod) ([]objects.KeyMetrics, []objects.KeyMetricsTTM, error) {
	// - [Key Metrics](https://site.financialmodelingprep.com/developer/docs#key-metrics-statement-analysis)
	// - [Key Metrics TTM](https://site.financialmodelingprep.com/developer/docs#key-metrics-ttm-statement-analysis)

	KeyMetrics, err := fmp.APIClient.CompanyValuation.KeyMetrics(objects.RequestKeyMetrics{
		Symbol: Symbol,
		Period: Period,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error getting key metrics: %s", err)
	}

	KeyMetricsTTM, err := fmp.APIClient.CompanyValuation.KeyMetricsTTM(Symbol)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting key metrics TTM: %s", err)
	}

	return KeyMetrics, KeyMetricsTTM, nil
}

func (fmp *FMPAPI) GetRatios(Symbol string, Period objects.CompanyValuationPeriod) ([]objects.FinancialRatios, []objects.FinancialRatiosTTM, error) {
	// - [Ratios](https://site.financialmodelingprep.com/developer/docs#ratios-statement-analysis)
	// - [Ratios TTM](https://site.financialmodelingprep.com/developer/docs#ratios-ttm-statement-analysis)
	Ratios, err := fmp.APIClient.CompanyValuation.FinancialRatios(objects.RequestFinancialRatios{
		Symbol: Symbol,
		Period: Period,
	})

	if err != nil {
		return nil, nil, fmt.Errorf("error getting financial ratios: %s", err)
	}

	RatiosTTM, err := fmp.APIClient.CompanyValuation.FinancialRatiosTTM(Symbol)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting financial ratios TTM: %s", err)
	}

	return Ratios, RatiosTTM, nil
}

func (fmp *FMPAPI) GetCashFlowStatementGrowth(Symbol string, Period objects.CompanyValuationPeriod) ([]objects.CashFlowStatementGrowth, error) {
	// - [Cashflow Growth](https://site.financialmodelingprep.com/developer/docs#cashflow-growth-statement-analysis)
	CashFlowGrowth, err := fmp.APIClient.CompanyValuation.CashFlowStatementGrowth(objects.RequestCashFlowStatementGrowth{
		Symbol: Symbol,
		Period: Period,
	})

	if err != nil {
		return nil, fmt.Errorf("error getting cash flow growth: %s", err)
	}

	return CashFlowGrowth, nil
}

func (fmp *FMPAPI) GetIncomeStatementGrowth(Symbol string, Period objects.CompanyValuationPeriod) ([]objects.IncomeStatementGrowth, error) {
	// - [Income Growth](https://site.financialmodelingprep.com/developer/docs#income-growth-statement-analysis)
	IncomeGrowth, err := fmp.APIClient.CompanyValuation.IncomeStatementGrowth(objects.RequestIncomeStatementGrowth{
		Symbol: Symbol,
		Period: Period,
	})

	if err != nil {
		return nil, fmt.Errorf("error getting income growth: %s", err)
	}

	return IncomeGrowth, nil
}

func (fmp *FMPAPI) GetBalanceSheetStatementGrowth(Symbol string, Period objects.CompanyValuationPeriod) ([]objects.BalanceSheetStatementGrowth, error) {
	// - [Balance Sheet Growth](https://site.financialmodelingprep.com/developer/docs#balance-sheet-growth-statement-analysis)
	BalanceSheetGrowth, err := fmp.APIClient.CompanyValuation.BalanceSheetStatementGrowth(objects.RequestBalanceSheetStatementGrowth{
		Symbol: Symbol,
		Period: Period,
	})

	if err != nil {
		return nil, fmt.Errorf("error getting balance sheet growth: %s", err)
	}

	return BalanceSheetGrowth, nil
}

func (fmp *FMPAPI) GetFinancialStatementGrowth(Symbol string, Period objects.CompanyValuationPeriod) ([]objects.FinancialStatementsGrowth, error) {
	// - [Financial Growth](https://site.financialmodelingprep.com/developer/docs#financial-growth-statement-analysis)
	FinancialGrowth, err := fmp.APIClient.CompanyValuation.FinancialStatementsGrowth(objects.RequestFinancialStatementsGrowth{
		Symbol: Symbol,
		Period: Period,
	})

	if err != nil {
		return nil, fmt.Errorf("error getting financial growth: %s", err)
	}

	return FinancialGrowth, nil
}
