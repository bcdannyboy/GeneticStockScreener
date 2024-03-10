package FMP

import "github.com/spacecodewor/fmpcloud-go/objects"

func GetAverageKeyMetrics(KeyMetrics []objects.KeyMetrics) objects.KeyMetrics {
	KeyMetricsAvg := objects.KeyMetrics{}

	for _, Metric := range KeyMetrics {
		KeyMetricsAvg.RevenuePerShare += Metric.RevenuePerShare
		KeyMetricsAvg.NetIncomePerShare += Metric.NetIncomePerShare
		KeyMetricsAvg.OperatingCashFlowPerShare += Metric.OperatingCashFlowPerShare
		KeyMetricsAvg.FreeCashFlowPerShare += Metric.FreeCashFlowPerShare
		KeyMetricsAvg.CashPerShare += Metric.CashPerShare
		KeyMetricsAvg.BookValuePerShare += Metric.BookValuePerShare
		KeyMetricsAvg.TangibleBookValuePerShare += Metric.TangibleBookValuePerShare
		KeyMetricsAvg.ShareholdersEquityPerShare += Metric.ShareholdersEquityPerShare
		KeyMetricsAvg.InterestDebtPerShare += Metric.InterestDebtPerShare
		KeyMetricsAvg.MarketCap += Metric.MarketCap
		KeyMetricsAvg.EnterpriseValue += Metric.EnterpriseValue
		KeyMetricsAvg.PeRatio += Metric.PeRatio
		KeyMetricsAvg.PriceToSalesRatio += Metric.PriceToSalesRatio
		KeyMetricsAvg.Pocfratio += Metric.Pocfratio
		KeyMetricsAvg.PfcfRatio += Metric.PfcfRatio
		KeyMetricsAvg.PbRatio += Metric.PbRatio
		KeyMetricsAvg.PtbRatio += Metric.PtbRatio
		KeyMetricsAvg.EvToSales += Metric.EvToSales
		KeyMetricsAvg.EnterpriseValueOverEBITDA += Metric.EnterpriseValueOverEBITDA
		KeyMetricsAvg.EvToOperatingCashFlow += Metric.EvToOperatingCashFlow
		KeyMetricsAvg.EvToFreeCashFlow += Metric.EvToFreeCashFlow
		KeyMetricsAvg.EarningsYield += Metric.EarningsYield
		KeyMetricsAvg.FreeCashFlowYield += Metric.FreeCashFlowYield
		KeyMetricsAvg.DebtToEquity += Metric.DebtToEquity
		KeyMetricsAvg.DebtToAssets += Metric.DebtToAssets
		KeyMetricsAvg.NetDebtToEBITDA += Metric.NetDebtToEBITDA
		KeyMetricsAvg.CurrentRatio += Metric.CurrentRatio
		KeyMetricsAvg.InterestCoverage += Metric.InterestCoverage
		KeyMetricsAvg.IncomeQuality += Metric.IncomeQuality
		KeyMetricsAvg.DividendYield += Metric.DividendYield
		KeyMetricsAvg.PayoutRatio += Metric.PayoutRatio
		KeyMetricsAvg.SalesGeneralAndAdministrativeToRevenue += Metric.SalesGeneralAndAdministrativeToRevenue
		KeyMetricsAvg.ResearchAndDdevelopementToRevenue += Metric.ResearchAndDdevelopementToRevenue
		KeyMetricsAvg.IntangiblesToTotalAssets += Metric.IntangiblesToTotalAssets
		KeyMetricsAvg.CapexToOperatingCashFlow += Metric.CapexToOperatingCashFlow
		KeyMetricsAvg.CapexToRevenue += Metric.CapexToRevenue
		KeyMetricsAvg.CapexToDepreciation += Metric.CapexToDepreciation
		KeyMetricsAvg.StockBasedCompensationToRevenue += Metric.StockBasedCompensationToRevenue
		KeyMetricsAvg.GrahamNumber += Metric.GrahamNumber
		KeyMetricsAvg.Roic += Metric.Roic
		KeyMetricsAvg.ReturnOnTangibleAssets += Metric.ReturnOnTangibleAssets
		KeyMetricsAvg.GrahamNetNet += Metric.GrahamNetNet
		KeyMetricsAvg.WorkingCapital += Metric.WorkingCapital
		KeyMetricsAvg.TangibleAssetValue += Metric.TangibleAssetValue
		KeyMetricsAvg.NetCurrentAssetValue += Metric.NetCurrentAssetValue
		KeyMetricsAvg.InvestedCapital += Metric.InvestedCapital
		KeyMetricsAvg.AverageReceivables += Metric.AverageReceivables
		KeyMetricsAvg.AveragePayables += Metric.AveragePayables
		KeyMetricsAvg.AverageInventory += Metric.AverageInventory
		KeyMetricsAvg.DaysSalesOutstanding += Metric.DaysSalesOutstanding
		KeyMetricsAvg.DaysPayablesOutstanding += Metric.DaysPayablesOutstanding
		KeyMetricsAvg.DaysOfInventoryOnHand += Metric.DaysOfInventoryOnHand
		KeyMetricsAvg.ReceivablesTurnover += Metric.ReceivablesTurnover
		KeyMetricsAvg.PayablesTurnover += Metric.PayablesTurnover
		KeyMetricsAvg.InventoryTurnover += Metric.InventoryTurnover
		KeyMetricsAvg.Roe += Metric.Roe
		KeyMetricsAvg.CapexPerShare += Metric.CapexPerShare
	}

	KeyMetricsAvg.RevenuePerShare = KeyMetricsAvg.RevenuePerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.NetIncomePerShare = KeyMetricsAvg.NetIncomePerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.OperatingCashFlowPerShare = KeyMetricsAvg.OperatingCashFlowPerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.FreeCashFlowPerShare = KeyMetricsAvg.FreeCashFlowPerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.CashPerShare = KeyMetricsAvg.CashPerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.BookValuePerShare = KeyMetricsAvg.BookValuePerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.TangibleBookValuePerShare = KeyMetricsAvg.TangibleBookValuePerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.ShareholdersEquityPerShare = KeyMetricsAvg.ShareholdersEquityPerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.InterestDebtPerShare = KeyMetricsAvg.InterestDebtPerShare / float64(len(KeyMetrics))
	KeyMetricsAvg.MarketCap = KeyMetricsAvg.MarketCap / float64(len(KeyMetrics))
	KeyMetricsAvg.EnterpriseValue = KeyMetricsAvg.EnterpriseValue / float64(len(KeyMetrics))
	KeyMetricsAvg.PeRatio = KeyMetricsAvg.PeRatio / float64(len(KeyMetrics))
	KeyMetricsAvg.PriceToSalesRatio = KeyMetricsAvg.PriceToSalesRatio / float64(len(KeyMetrics))
	KeyMetricsAvg.Pocfratio = KeyMetricsAvg.Pocfratio / float64(len(KeyMetrics))
	KeyMetricsAvg.PfcfRatio = KeyMetricsAvg.PfcfRatio / float64(len(KeyMetrics))
	KeyMetricsAvg.PbRatio = KeyMetricsAvg.PbRatio / float64(len(KeyMetrics))
	KeyMetricsAvg.PtbRatio = KeyMetricsAvg.PtbRatio / float64(len(KeyMetrics))
	KeyMetricsAvg.EvToSales = KeyMetricsAvg.EvToSales / float64(len(KeyMetrics))
	KeyMetricsAvg.EnterpriseValueOverEBITDA = KeyMetricsAvg.EnterpriseValueOverEBITDA / float64(len(KeyMetrics))
	KeyMetricsAvg.EvToOperatingCashFlow = KeyMetricsAvg.EvToOperatingCashFlow / float64(len(KeyMetrics))
	KeyMetricsAvg.EvToFreeCashFlow = KeyMetricsAvg.EvToFreeCashFlow / float64(len(KeyMetrics))
	KeyMetricsAvg.EarningsYield = KeyMetricsAvg.EarningsYield / float64(len(KeyMetrics))
	KeyMetricsAvg.FreeCashFlowYield = KeyMetricsAvg.FreeCashFlowYield / float64(len(KeyMetrics))
	KeyMetricsAvg.DebtToEquity = KeyMetricsAvg.DebtToEquity / float64(len(KeyMetrics))
	KeyMetricsAvg.DebtToAssets = KeyMetricsAvg.DebtToAssets / float64(len(KeyMetrics))
	KeyMetricsAvg.NetDebtToEBITDA = KeyMetricsAvg.NetDebtToEBITDA / float64(len(KeyMetrics))
	KeyMetricsAvg.CurrentRatio = KeyMetricsAvg.CurrentRatio / float64(len(KeyMetrics))
	KeyMetricsAvg.InterestCoverage = KeyMetricsAvg.InterestCoverage / float64(len(KeyMetrics))
	KeyMetricsAvg.IncomeQuality = KeyMetricsAvg.IncomeQuality / float64(len(KeyMetrics))
	KeyMetricsAvg.DividendYield = KeyMetricsAvg.DividendYield / float64(len(KeyMetrics))
	KeyMetricsAvg.PayoutRatio = KeyMetricsAvg.PayoutRatio / float64(len(KeyMetrics))
	KeyMetricsAvg.SalesGeneralAndAdministrativeToRevenue = KeyMetricsAvg.SalesGeneralAndAdministrativeToRevenue / float64(len(KeyMetrics))
	KeyMetricsAvg.ResearchAndDdevelopementToRevenue = KeyMetricsAvg.ResearchAndDdevelopementToRevenue / float64(len(KeyMetrics))
	KeyMetricsAvg.IntangiblesToTotalAssets = KeyMetricsAvg.IntangiblesToTotalAssets / float64(len(KeyMetrics))
	KeyMetricsAvg.CapexToOperatingCashFlow = KeyMetricsAvg.CapexToOperatingCashFlow / float64(len(KeyMetrics))
	KeyMetricsAvg.CapexToRevenue = KeyMetricsAvg.CapexToRevenue / float64(len(KeyMetrics))
	KeyMetricsAvg.CapexToDepreciation = KeyMetricsAvg.CapexToDepreciation / float64(len(KeyMetrics))
	KeyMetricsAvg.StockBasedCompensationToRevenue = KeyMetricsAvg.StockBasedCompensationToRevenue / float64(len(KeyMetrics))
	KeyMetricsAvg.GrahamNumber = KeyMetricsAvg.GrahamNumber / float64(len(KeyMetrics))
	KeyMetricsAvg.Roic = KeyMetricsAvg.Roic / float64(len(KeyMetrics))
	KeyMetricsAvg.ReturnOnTangibleAssets = KeyMetricsAvg.ReturnOnTangibleAssets / float64(len(KeyMetrics))
	KeyMetricsAvg.GrahamNetNet = KeyMetricsAvg.GrahamNetNet / float64(len(KeyMetrics))
	KeyMetricsAvg.WorkingCapital = KeyMetricsAvg.WorkingCapital / float64(len(KeyMetrics))
	KeyMetricsAvg.TangibleAssetValue = KeyMetricsAvg.TangibleAssetValue / float64(len(KeyMetrics))
	KeyMetricsAvg.NetCurrentAssetValue = KeyMetricsAvg.NetCurrentAssetValue / float64(len(KeyMetrics))
	KeyMetricsAvg.InvestedCapital = KeyMetricsAvg.InvestedCapital / float64(len(KeyMetrics))
	KeyMetricsAvg.AverageReceivables = KeyMetricsAvg.AverageReceivables / float64(len(KeyMetrics))
	KeyMetricsAvg.AveragePayables = KeyMetricsAvg.AveragePayables / float64(len(KeyMetrics))
	KeyMetricsAvg.AverageInventory = KeyMetricsAvg.AverageInventory / float64(len(KeyMetrics))
	KeyMetricsAvg.DaysSalesOutstanding = KeyMetricsAvg.DaysSalesOutstanding / float64(len(KeyMetrics))
	KeyMetricsAvg.DaysPayablesOutstanding = KeyMetricsAvg.DaysPayablesOutstanding / float64(len(KeyMetrics))
	KeyMetricsAvg.DaysOfInventoryOnHand = KeyMetricsAvg.DaysOfInventoryOnHand / float64(len(KeyMetrics))
	KeyMetricsAvg.ReceivablesTurnover = KeyMetricsAvg.ReceivablesTurnover / float64(len(KeyMetrics))
	KeyMetricsAvg.PayablesTurnover = KeyMetricsAvg.PayablesTurnover / float64(len(KeyMetrics))
	KeyMetricsAvg.InventoryTurnover = KeyMetricsAvg.InventoryTurnover / float64(len(KeyMetrics))
	KeyMetricsAvg.Roe = KeyMetricsAvg.Roe / float64(len(KeyMetrics))
	KeyMetricsAvg.CapexPerShare = KeyMetricsAvg.CapexPerShare / float64(len(KeyMetrics))

	return KeyMetricsAvg
}

func GetAverageFinancialRatios(Ratios []objects.FinancialRatios) objects.FinancialRatios {
	RatiosAvg := objects.FinancialRatios{}

	for _, Ratio := range Ratios {
		RatiosAvg.CurrentRatio += Ratio.CurrentRatio
		RatiosAvg.QuickRatio += Ratio.QuickRatio
		RatiosAvg.CashRatio += Ratio.CashRatio
		RatiosAvg.DaysOfSalesOutstanding += Ratio.DaysOfSalesOutstanding
		RatiosAvg.DaysOfInventoryOutstanding += Ratio.DaysOfInventoryOutstanding
		RatiosAvg.OperatingCycle += Ratio.OperatingCycle
		RatiosAvg.DaysOfPayablesOutstanding += Ratio.DaysOfPayablesOutstanding
		RatiosAvg.CashConversionCycle += Ratio.CashConversionCycle
		RatiosAvg.GrossProfitMargin += Ratio.GrossProfitMargin
		RatiosAvg.OperatingProfitMargin += Ratio.OperatingProfitMargin
		RatiosAvg.PretaxProfitMargin += Ratio.PretaxProfitMargin
		RatiosAvg.NetProfitMargin += Ratio.NetProfitMargin
		RatiosAvg.EffectiveTaxRate += Ratio.EffectiveTaxRate
		RatiosAvg.ReturnOnAssets += Ratio.ReturnOnAssets
		RatiosAvg.ReturnOnEquity += Ratio.ReturnOnEquity
		RatiosAvg.ReturnOnCapitalEmployed += Ratio.ReturnOnCapitalEmployed
		RatiosAvg.NetIncomePerEBT += Ratio.NetIncomePerEBT
		RatiosAvg.EbtPerEbit += Ratio.EbtPerEbit
		RatiosAvg.EbitPerRevenue += Ratio.EbitPerRevenue
		RatiosAvg.DebtRatio += Ratio.DebtRatio
		RatiosAvg.DebtEquityRatio += Ratio.DebtEquityRatio
		RatiosAvg.LongTermDebtToCapitalization += Ratio.LongTermDebtToCapitalization
		RatiosAvg.TotalDebtToCapitalization += Ratio.TotalDebtToCapitalization
		RatiosAvg.InterestCoverage += Ratio.InterestCoverage
		RatiosAvg.CashFlowToDebtRatio += Ratio.CashFlowToDebtRatio
		RatiosAvg.CompanyEquityMultiplier += Ratio.CompanyEquityMultiplier
		RatiosAvg.ReceivablesTurnover += Ratio.ReceivablesTurnover
		RatiosAvg.PayablesTurnover += Ratio.PayablesTurnover
		RatiosAvg.InventoryTurnover += Ratio.InventoryTurnover
		RatiosAvg.FixedAssetTurnover += Ratio.FixedAssetTurnover
		RatiosAvg.AssetTurnover += Ratio.AssetTurnover
		RatiosAvg.OperatingCashFlowPerShare += Ratio.OperatingCashFlowPerShare
		RatiosAvg.FreeCashFlowPerShare += Ratio.FreeCashFlowPerShare
		RatiosAvg.CashPerShare += Ratio.CashPerShare
		RatiosAvg.PayoutRatio += Ratio.PayoutRatio
		RatiosAvg.OperatingCashFlowSalesRatio += Ratio.OperatingCashFlowSalesRatio
		RatiosAvg.FreeCashFlowOperatingCashFlowRatio += Ratio.FreeCashFlowOperatingCashFlowRatio
		RatiosAvg.CashFlowCoverageRatios += Ratio.CashFlowCoverageRatios
		RatiosAvg.ShortTermCoverageRatios += Ratio.ShortTermCoverageRatios
		RatiosAvg.CapitalExpenditureCoverageRatio += Ratio.CapitalExpenditureCoverageRatio
		RatiosAvg.DividendPaidAndCapexCoverageRatio += Ratio.DividendPaidAndCapexCoverageRatio
		RatiosAvg.DividendPayoutRatio += Ratio.DividendPayoutRatio
		RatiosAvg.PriceBookValueRatio += Ratio.PriceBookValueRatio
		RatiosAvg.PriceToBookRatio += Ratio.PriceToBookRatio
		RatiosAvg.PriceToSalesRatio += Ratio.PriceToSalesRatio
		RatiosAvg.PriceEarningsRatio += Ratio.PriceEarningsRatio
		RatiosAvg.PriceToFreeCashFlowsRatio += Ratio.PriceToFreeCashFlowsRatio
		RatiosAvg.PriceToOperatingCashFlowsRatio += Ratio.PriceToOperatingCashFlowsRatio
		RatiosAvg.PriceCashFlowRatio += Ratio.PriceCashFlowRatio
		RatiosAvg.PriceEarningsToGrowthRatio += Ratio.PriceEarningsToGrowthRatio
		RatiosAvg.PriceSalesRatio += Ratio.PriceSalesRatio
		RatiosAvg.DividendYield += Ratio.DividendYield
		RatiosAvg.EnterpriseValueMultiple += Ratio.EnterpriseValueMultiple
		RatiosAvg.PriceFairValue += Ratio.PriceFairValue
	}

	RatiosAvg.CurrentRatio = RatiosAvg.CurrentRatio / float64(len(Ratios))
	RatiosAvg.QuickRatio = RatiosAvg.QuickRatio / float64(len(Ratios))
	RatiosAvg.CashRatio = RatiosAvg.CashRatio / float64(len(Ratios))
	RatiosAvg.DaysOfSalesOutstanding = RatiosAvg.DaysOfSalesOutstanding / float64(len(Ratios))
	RatiosAvg.DaysOfInventoryOutstanding = RatiosAvg.DaysOfInventoryOutstanding / float64(len(Ratios))
	RatiosAvg.OperatingCycle = RatiosAvg.OperatingCycle / float64(len(Ratios))
	RatiosAvg.DaysOfPayablesOutstanding = RatiosAvg.DaysOfPayablesOutstanding / float64(len(Ratios))
	RatiosAvg.CashConversionCycle = RatiosAvg.CashConversionCycle / float64(len(Ratios))
	RatiosAvg.GrossProfitMargin = RatiosAvg.GrossProfitMargin / float64(len(Ratios))
	RatiosAvg.OperatingProfitMargin = RatiosAvg.OperatingProfitMargin / float64(len(Ratios))
	RatiosAvg.PretaxProfitMargin = RatiosAvg.PretaxProfitMargin / float64(len(Ratios))
	RatiosAvg.NetProfitMargin = RatiosAvg.NetProfitMargin / float64(len(Ratios))
	RatiosAvg.EffectiveTaxRate = RatiosAvg.EffectiveTaxRate / float64(len(Ratios))
	RatiosAvg.ReturnOnAssets = RatiosAvg.ReturnOnAssets / float64(len(Ratios))
	RatiosAvg.ReturnOnEquity = RatiosAvg.ReturnOnEquity / float64(len(Ratios))
	RatiosAvg.ReturnOnCapitalEmployed = RatiosAvg.ReturnOnCapitalEmployed / float64(len(Ratios))
	RatiosAvg.NetIncomePerEBT = RatiosAvg.NetIncomePerEBT / float64(len(Ratios))
	RatiosAvg.EbtPerEbit = RatiosAvg.EbtPerEbit / float64(len(Ratios))
	RatiosAvg.EbitPerRevenue = RatiosAvg.EbitPerRevenue / float64(len(Ratios))
	RatiosAvg.DebtRatio = RatiosAvg.DebtRatio / float64(len(Ratios))
	RatiosAvg.DebtEquityRatio = RatiosAvg.DebtEquityRatio / float64(len(Ratios))
	RatiosAvg.LongTermDebtToCapitalization = RatiosAvg.LongTermDebtToCapitalization / float64(len(Ratios))
	RatiosAvg.TotalDebtToCapitalization = RatiosAvg.TotalDebtToCapitalization / float64(len(Ratios))
	RatiosAvg.InterestCoverage = RatiosAvg.InterestCoverage / float64(len(Ratios))
	RatiosAvg.CashFlowToDebtRatio = RatiosAvg.CashFlowToDebtRatio / float64(len(Ratios))
	RatiosAvg.CompanyEquityMultiplier = RatiosAvg.CompanyEquityMultiplier / float64(len(Ratios))
	RatiosAvg.ReceivablesTurnover = RatiosAvg.ReceivablesTurnover / float64(len(Ratios))
	RatiosAvg.PayablesTurnover = RatiosAvg.PayablesTurnover / float64(len(Ratios))
	RatiosAvg.InventoryTurnover = RatiosAvg.InventoryTurnover / float64(len(Ratios))
	RatiosAvg.FixedAssetTurnover = RatiosAvg.FixedAssetTurnover / float64(len(Ratios))
	RatiosAvg.AssetTurnover = RatiosAvg.AssetTurnover / float64(len(Ratios))
	RatiosAvg.OperatingCashFlowPerShare = RatiosAvg.OperatingCashFlowPerShare / float64(len(Ratios))
	RatiosAvg.FreeCashFlowPerShare = RatiosAvg.FreeCashFlowPerShare / float64(len(Ratios))
	RatiosAvg.CashPerShare = RatiosAvg.CashPerShare / float64(len(Ratios))
	RatiosAvg.PayoutRatio = RatiosAvg.PayoutRatio / float64(len(Ratios))
	RatiosAvg.OperatingCashFlowSalesRatio = RatiosAvg.OperatingCashFlowSalesRatio / float64(len(Ratios))
	RatiosAvg.FreeCashFlowOperatingCashFlowRatio = RatiosAvg.FreeCashFlowOperatingCashFlowRatio / float64(len(Ratios))
	RatiosAvg.CashFlowCoverageRatios = RatiosAvg.CashFlowCoverageRatios / float64(len(Ratios))
	RatiosAvg.ShortTermCoverageRatios = RatiosAvg.ShortTermCoverageRatios / float64(len(Ratios))
	RatiosAvg.CapitalExpenditureCoverageRatio = RatiosAvg.CapitalExpenditureCoverageRatio / float64(len(Ratios))
	RatiosAvg.DividendPaidAndCapexCoverageRatio = RatiosAvg.DividendPaidAndCapexCoverageRatio / float64(len(Ratios))
	RatiosAvg.DividendPayoutRatio = RatiosAvg.DividendPayoutRatio / float64(len(Ratios))
	RatiosAvg.PriceBookValueRatio = RatiosAvg.PriceBookValueRatio / float64(len(Ratios))
	RatiosAvg.PriceToBookRatio = RatiosAvg.PriceToBookRatio / float64(len(Ratios))
	RatiosAvg.PriceToSalesRatio = RatiosAvg.PriceToSalesRatio / float64(len(Ratios))
	RatiosAvg.PriceEarningsRatio = RatiosAvg.PriceEarningsRatio / float64(len(Ratios))
	RatiosAvg.PriceToFreeCashFlowsRatio = RatiosAvg.PriceToFreeCashFlowsRatio / float64(len(Ratios))
	RatiosAvg.PriceToOperatingCashFlowsRatio = RatiosAvg.PriceToOperatingCashFlowsRatio / float64(len(Ratios))
	RatiosAvg.PriceCashFlowRatio = RatiosAvg.PriceCashFlowRatio / float64(len(Ratios))
	RatiosAvg.PriceEarningsToGrowthRatio = RatiosAvg.PriceEarningsToGrowthRatio / float64(len(Ratios))
	RatiosAvg.PriceSalesRatio = RatiosAvg.PriceSalesRatio / float64(len(Ratios))
	RatiosAvg.DividendYield = RatiosAvg.DividendYield / float64(len(Ratios))
	RatiosAvg.EnterpriseValueMultiple = RatiosAvg.EnterpriseValueMultiple / float64(len(Ratios))
	RatiosAvg.PriceFairValue = RatiosAvg.PriceFairValue / float64(len(Ratios))

	return RatiosAvg
}

func GetAverageCashFlowStatementGrowth(CFSG []objects.CashFlowStatementGrowth) objects.CashFlowStatementGrowth {
	AvgCFSG := objects.CashFlowStatementGrowth{}

	for _, Growth := range CFSG {
		AvgCFSG.GrowthNetIncome += Growth.GrowthNetIncome
		AvgCFSG.GrowthDepreciationAndAmortization += Growth.GrowthDepreciationAndAmortization
		AvgCFSG.GrowthDeferredIncomeTax += Growth.GrowthDeferredIncomeTax
		AvgCFSG.GrowthStockBasedCompensation += Growth.GrowthStockBasedCompensation
		AvgCFSG.GrowthChangeInWorkingCapital += Growth.GrowthChangeInWorkingCapital
		AvgCFSG.GrowthAccountsReceivables += Growth.GrowthAccountsReceivables
		AvgCFSG.GrowthInventory += Growth.GrowthInventory
		AvgCFSG.GrowthAccountsPayables += Growth.GrowthAccountsPayables
		AvgCFSG.GrowthOtherWorkingCapital += Growth.GrowthOtherWorkingCapital
		AvgCFSG.GrowthOtherNonCashItems += Growth.GrowthOtherNonCashItems
		AvgCFSG.GrowthNetCashProvidedByOperatingActivites += Growth.GrowthNetCashProvidedByOperatingActivites
		AvgCFSG.GrowthInvestmentsInPropertyPlantAndEquipment += Growth.GrowthInvestmentsInPropertyPlantAndEquipment
		AvgCFSG.GrowthAcquisitionsNet += Growth.GrowthAcquisitionsNet
		AvgCFSG.GrowthPurchasesOfInvestments += Growth.GrowthPurchasesOfInvestments
		AvgCFSG.GrowthSalesMaturitiesOfInvestments += Growth.GrowthSalesMaturitiesOfInvestments
		AvgCFSG.GrowthOtherInvestingActivites += Growth.GrowthOtherInvestingActivites
		AvgCFSG.GrowthNetCashUsedForInvestingActivites += Growth.GrowthNetCashUsedForInvestingActivites
		AvgCFSG.GrowthDebtRepayment += Growth.GrowthDebtRepayment
		AvgCFSG.GrowthCommonStockIssued += Growth.GrowthCommonStockIssued
		AvgCFSG.GrowthCommonStockRepurchased += Growth.GrowthCommonStockRepurchased
		AvgCFSG.GrowthDividendsPaid += Growth.GrowthDividendsPaid
		AvgCFSG.GrowthOtherFinancingActivites += Growth.GrowthOtherFinancingActivites
		AvgCFSG.GrowthNetCashUsedProvidedByFinancingActivities += Growth.GrowthNetCashUsedProvidedByFinancingActivities
		AvgCFSG.GrowthEffectOfForexChangesOnCash += Growth.GrowthEffectOfForexChangesOnCash
		AvgCFSG.GrowthNetChangeInCash += Growth.GrowthNetChangeInCash
		AvgCFSG.GrowthCashAtEndOfPeriod += Growth.GrowthCashAtEndOfPeriod
		AvgCFSG.GrowthCashAtBeginningOfPeriod += Growth.GrowthCashAtBeginningOfPeriod
		AvgCFSG.GrowthOperatingCashFlow += Growth.GrowthOperatingCashFlow
		AvgCFSG.GrowthCapitalExpenditure += Growth.GrowthCapitalExpenditure
		AvgCFSG.GrowthFreeCashFlow += Growth.GrowthFreeCashFlow
	}

	AvgCFSG.GrowthNetIncome = AvgCFSG.GrowthNetIncome / float64(len(CFSG))
	AvgCFSG.GrowthDepreciationAndAmortization = AvgCFSG.GrowthDepreciationAndAmortization / float64(len(CFSG))
	AvgCFSG.GrowthDeferredIncomeTax = AvgCFSG.GrowthDeferredIncomeTax / float64(len(CFSG))
	AvgCFSG.GrowthStockBasedCompensation = AvgCFSG.GrowthStockBasedCompensation / float64(len(CFSG))
	AvgCFSG.GrowthChangeInWorkingCapital = AvgCFSG.GrowthChangeInWorkingCapital / float64(len(CFSG))
	AvgCFSG.GrowthAccountsReceivables = AvgCFSG.GrowthAccountsReceivables / float64(len(CFSG))
	AvgCFSG.GrowthInventory = AvgCFSG.GrowthInventory / float64(len(CFSG))
	AvgCFSG.GrowthAccountsPayables = AvgCFSG.GrowthAccountsPayables / float64(len(CFSG))
	AvgCFSG.GrowthOtherWorkingCapital = AvgCFSG.GrowthOtherWorkingCapital / float64(len(CFSG))
	AvgCFSG.GrowthOtherNonCashItems = AvgCFSG.GrowthOtherNonCashItems / float64(len(CFSG))
	AvgCFSG.GrowthNetCashProvidedByOperatingActivites = AvgCFSG.GrowthNetCashProvidedByOperatingActivites / float64(len(CFSG))
	AvgCFSG.GrowthInvestmentsInPropertyPlantAndEquipment = AvgCFSG.GrowthInvestmentsInPropertyPlantAndEquipment / float64(len(CFSG))
	AvgCFSG.GrowthAcquisitionsNet = AvgCFSG.GrowthAcquisitionsNet / float64(len(CFSG))
	AvgCFSG.GrowthPurchasesOfInvestments = AvgCFSG.GrowthPurchasesOfInvestments / float64(len(CFSG))
	AvgCFSG.GrowthSalesMaturitiesOfInvestments = AvgCFSG.GrowthSalesMaturitiesOfInvestments / float64(len(CFSG))
	AvgCFSG.GrowthOtherInvestingActivites = AvgCFSG.GrowthOtherInvestingActivites / float64(len(CFSG))
	AvgCFSG.GrowthNetCashUsedForInvestingActivites = AvgCFSG.GrowthNetCashUsedForInvestingActivites / float64(len(CFSG))
	AvgCFSG.GrowthDebtRepayment = AvgCFSG.GrowthDebtRepayment / float64(len(CFSG))
	AvgCFSG.GrowthCommonStockIssued = AvgCFSG.GrowthCommonStockIssued / float64(len(CFSG))
	AvgCFSG.GrowthCommonStockRepurchased = AvgCFSG.GrowthCommonStockRepurchased / float64(len(CFSG))
	AvgCFSG.GrowthDividendsPaid = AvgCFSG.GrowthDividendsPaid / float64(len(CFSG))
	AvgCFSG.GrowthOtherFinancingActivites = AvgCFSG.GrowthOtherFinancingActivites / float64(len(CFSG))
	AvgCFSG.GrowthNetCashUsedProvidedByFinancingActivities = AvgCFSG.GrowthNetCashUsedProvidedByFinancingActivities / float64(len(CFSG))
	AvgCFSG.GrowthEffectOfForexChangesOnCash = AvgCFSG.GrowthEffectOfForexChangesOnCash / float64(len(CFSG))
	AvgCFSG.GrowthNetChangeInCash = AvgCFSG.GrowthNetChangeInCash / float64(len(CFSG))
	AvgCFSG.GrowthCashAtEndOfPeriod = AvgCFSG.GrowthCashAtEndOfPeriod / float64(len(CFSG))
	AvgCFSG.GrowthCashAtBeginningOfPeriod = AvgCFSG.GrowthCashAtBeginningOfPeriod / float64(len(CFSG))
	AvgCFSG.GrowthOperatingCashFlow = AvgCFSG.GrowthOperatingCashFlow / float64(len(CFSG))
	AvgCFSG.GrowthCapitalExpenditure = AvgCFSG.GrowthCapitalExpenditure / float64(len(CFSG))
	AvgCFSG.GrowthFreeCashFlow = AvgCFSG.GrowthFreeCashFlow / float64(len(CFSG))

	return AvgCFSG
}

func GetAverageIncomeStatementGrowth(ISG []objects.IncomeStatementGrowth) objects.IncomeStatementGrowth {
	AvgISG := objects.IncomeStatementGrowth{}

	for _, Growth := range ISG {
		AvgISG.GrowthRevenue += Growth.GrowthRevenue
		AvgISG.GrowthCostOfRevenue += Growth.GrowthCostOfRevenue
		AvgISG.GrowthGrossProfit += Growth.GrowthGrossProfit
		AvgISG.GrowthGrossProfitRatio += Growth.GrowthGrossProfitRatio
		AvgISG.GrowthResearchAndDevelopmentExpenses += Growth.GrowthResearchAndDevelopmentExpenses
		AvgISG.GrowthGeneralAndAdministrativeExpenses += Growth.GrowthGeneralAndAdministrativeExpenses
		AvgISG.GrowthSellingAndMarketingExpenses += Growth.GrowthSellingAndMarketingExpenses
		AvgISG.GrowthOtherExpenses += Growth.GrowthOtherExpenses
		AvgISG.GrowthOperatingExpenses += Growth.GrowthOperatingExpenses
		AvgISG.GrowthCostAndExpenses += Growth.GrowthCostAndExpenses
		AvgISG.GrowthInterestExpense += Growth.GrowthInterestExpense
		AvgISG.GrowthDepreciationAndAmortization += Growth.GrowthDepreciationAndAmortization
		AvgISG.GrowthEBITDA += Growth.GrowthEBITDA
		AvgISG.GrowthEBITDARatio += Growth.GrowthEBITDARatio
		AvgISG.GrowthOperatingIncome += Growth.GrowthOperatingIncome
		AvgISG.GrowthOperatingIncomeRatio += Growth.GrowthOperatingIncomeRatio
		AvgISG.GrowthTotalOtherIncomeExpensesNet += Growth.GrowthTotalOtherIncomeExpensesNet
		AvgISG.GrowthIncomeBeforeTax += Growth.GrowthIncomeBeforeTax
		AvgISG.GrowthIncomeBeforeTaxRatio += Growth.GrowthIncomeBeforeTaxRatio
		AvgISG.GrowthIncomeTaxExpense += Growth.GrowthIncomeTaxExpense
		AvgISG.GrowthNetIncome += Growth.GrowthNetIncome
		AvgISG.GrowthNetIncomeRatio += Growth.GrowthNetIncomeRatio
		AvgISG.GrowthEPS += Growth.GrowthEPS
		AvgISG.GrowthEPSDiluted += Growth.GrowthEPSDiluted
		AvgISG.GrowthWeightedAverageShsOut += Growth.GrowthWeightedAverageShsOut
		AvgISG.GrowthWeightedAverageShsOutDil += Growth.GrowthWeightedAverageShsOutDil
	}

	AvgISG.GrowthRevenue = AvgISG.GrowthRevenue / float64(len(ISG))
	AvgISG.GrowthCostOfRevenue = AvgISG.GrowthCostOfRevenue / float64(len(ISG))
	AvgISG.GrowthGrossProfit = AvgISG.GrowthGrossProfit / float64(len(ISG))
	AvgISG.GrowthGrossProfitRatio = AvgISG.GrowthGrossProfitRatio / float64(len(ISG))
	AvgISG.GrowthResearchAndDevelopmentExpenses = AvgISG.GrowthResearchAndDevelopmentExpenses / float64(len(ISG))
	AvgISG.GrowthGeneralAndAdministrativeExpenses = AvgISG.GrowthGeneralAndAdministrativeExpenses / float64(len(ISG))
	AvgISG.GrowthSellingAndMarketingExpenses = AvgISG.GrowthSellingAndMarketingExpenses / float64(len(ISG))
	AvgISG.GrowthOtherExpenses = AvgISG.GrowthOtherExpenses / float64(len(ISG))
	AvgISG.GrowthOperatingExpenses = AvgISG.GrowthOperatingExpenses / float64(len(ISG))
	AvgISG.GrowthCostAndExpenses = AvgISG.GrowthCostAndExpenses / float64(len(ISG))
	AvgISG.GrowthInterestExpense = AvgISG.GrowthInterestExpense / float64(len(ISG))
	AvgISG.GrowthDepreciationAndAmortization = AvgISG.GrowthDepreciationAndAmortization / float64(len(ISG))
	AvgISG.GrowthEBITDA = AvgISG.GrowthEBITDA / float64(len(ISG))
	AvgISG.GrowthEBITDARatio = AvgISG.GrowthEBITDARatio / float64(len(ISG))
	AvgISG.GrowthOperatingIncome = AvgISG.GrowthOperatingIncome / float64(len(ISG))
	AvgISG.GrowthOperatingIncomeRatio = AvgISG.GrowthOperatingIncomeRatio / float64(len(ISG))
	AvgISG.GrowthTotalOtherIncomeExpensesNet = AvgISG.GrowthTotalOtherIncomeExpensesNet / float64(len(ISG))
	AvgISG.GrowthIncomeBeforeTax = AvgISG.GrowthIncomeBeforeTax / float64(len(ISG))
	AvgISG.GrowthIncomeBeforeTaxRatio = AvgISG.GrowthIncomeBeforeTaxRatio / float64(len(ISG))
	AvgISG.GrowthIncomeTaxExpense = AvgISG.GrowthIncomeTaxExpense / float64(len(ISG))
	AvgISG.GrowthNetIncome = AvgISG.GrowthNetIncome / float64(len(ISG))
	AvgISG.GrowthNetIncomeRatio = AvgISG.GrowthNetIncomeRatio / float64(len(ISG))
	AvgISG.GrowthEPS = AvgISG.GrowthEPS / float64(len(ISG))
	AvgISG.GrowthEPSDiluted = AvgISG.GrowthEPSDiluted / float64(len(ISG))
	AvgISG.GrowthWeightedAverageShsOut = AvgISG.GrowthWeightedAverageShsOut / float64(len(ISG))
	AvgISG.GrowthWeightedAverageShsOutDil = AvgISG.GrowthWeightedAverageShsOutDil / float64(len(ISG))

	return AvgISG
}

func GetAverageBalanceSheetStatementGrowth(BSSG []objects.BalanceSheetStatementGrowth) objects.BalanceSheetStatementGrowth {
	AvgBSSG := objects.BalanceSheetStatementGrowth{}

	for _, Growth := range BSSG {
		AvgBSSG.GrowthCashAndCashEquivalents += Growth.GrowthCashAndCashEquivalents
		AvgBSSG.GrowthShortTermInvestments += Growth.GrowthShortTermInvestments
		AvgBSSG.GrowthCashAndShortTermInvestments += Growth.GrowthCashAndShortTermInvestments
		AvgBSSG.GrowthNetReceivables += Growth.GrowthNetReceivables
		AvgBSSG.GrowthInventory += Growth.GrowthInventory
		AvgBSSG.GrowthOtherCurrentAssets += Growth.GrowthOtherCurrentAssets
		AvgBSSG.GrowthTotalCurrentAssets += Growth.GrowthTotalCurrentAssets
		AvgBSSG.GrowthPropertyPlantEquipmentNet += Growth.GrowthPropertyPlantEquipmentNet
		AvgBSSG.GrowthGoodwill += Growth.GrowthGoodwill
		AvgBSSG.GrowthIntangibleAssets += Growth.GrowthIntangibleAssets
		AvgBSSG.GrowthGoodwillAndIntangibleAssets += Growth.GrowthGoodwillAndIntangibleAssets
		AvgBSSG.GrowthLongTermInvestments += Growth.GrowthLongTermInvestments
		AvgBSSG.GrowthTaxAssets += Growth.GrowthTaxAssets
		AvgBSSG.GrowthOtherNonCurrentAssets += Growth.GrowthOtherNonCurrentAssets
		AvgBSSG.GrowthTotalNonCurrentAssets += Growth.GrowthTotalNonCurrentAssets
		AvgBSSG.GrowthOtherAssets += Growth.GrowthOtherAssets
		AvgBSSG.GrowthTotalAssets += Growth.GrowthTotalAssets
		AvgBSSG.GrowthAccountPayables += Growth.GrowthAccountPayables
		AvgBSSG.GrowthShortTermDebt += Growth.GrowthShortTermDebt
		AvgBSSG.GrowthTaxPayables += Growth.GrowthTaxPayables
		AvgBSSG.GrowthDeferredRevenue += Growth.GrowthDeferredRevenue
		AvgBSSG.GrowthOtherCurrentLiabilities += Growth.GrowthOtherCurrentLiabilities
		AvgBSSG.GrowthTotalCurrentLiabilities += Growth.GrowthTotalCurrentLiabilities
		AvgBSSG.GrowthLongTermDebt += Growth.GrowthLongTermDebt
		AvgBSSG.GrowthDeferredRevenueNonCurrent += Growth.GrowthDeferredRevenueNonCurrent
		AvgBSSG.GrowthDeferrredTaxLiabilitiesNonCurrent += Growth.GrowthDeferrredTaxLiabilitiesNonCurrent
		AvgBSSG.GrowthOtherNonCurrentLiabilities += Growth.GrowthOtherNonCurrentLiabilities
		AvgBSSG.GrowthTotalNonCurrentLiabilities += Growth.GrowthTotalNonCurrentLiabilities
		AvgBSSG.GrowthOtherLiabilities += Growth.GrowthOtherLiabilities
		AvgBSSG.GrowthTotalLiabilities += Growth.GrowthTotalLiabilities
		AvgBSSG.GrowthCommonStock += Growth.GrowthCommonStock
		AvgBSSG.GrowthRetainedEarnings += Growth.GrowthRetainedEarnings
		AvgBSSG.GrowthAccumulatedOtherComprehensiveIncomeLoss += Growth.GrowthAccumulatedOtherComprehensiveIncomeLoss
		AvgBSSG.GrowthOthertotalStockholdersEquity += Growth.GrowthOthertotalStockholdersEquity
		AvgBSSG.GrowthTotalStockholdersEquity += Growth.GrowthTotalStockholdersEquity
		AvgBSSG.GrowthTotalLiabilitiesAndStockholdersEquity += Growth.GrowthTotalLiabilitiesAndStockholdersEquity
		AvgBSSG.GrowthTotalInvestments += Growth.GrowthTotalInvestments
		AvgBSSG.GrowthTotalDebt += Growth.GrowthTotalDebt
		AvgBSSG.GrowthNetDebt += Growth.GrowthNetDebt
	}

	AvgBSSG.GrowthCashAndCashEquivalents = AvgBSSG.GrowthCashAndCashEquivalents / float64(len(BSSG))
	AvgBSSG.GrowthShortTermInvestments = AvgBSSG.GrowthShortTermInvestments / float64(len(BSSG))
	AvgBSSG.GrowthCashAndShortTermInvestments = AvgBSSG.GrowthCashAndShortTermInvestments / float64(len(BSSG))
	AvgBSSG.GrowthNetReceivables = AvgBSSG.GrowthNetReceivables / float64(len(BSSG))
	AvgBSSG.GrowthInventory = AvgBSSG.GrowthInventory / float64(len(BSSG))
	AvgBSSG.GrowthOtherCurrentAssets = AvgBSSG.GrowthOtherCurrentAssets / float64(len(BSSG))
	AvgBSSG.GrowthTotalCurrentAssets = AvgBSSG.GrowthTotalCurrentAssets / float64(len(BSSG))
	AvgBSSG.GrowthPropertyPlantEquipmentNet = AvgBSSG.GrowthPropertyPlantEquipmentNet / float64(len(BSSG))
	AvgBSSG.GrowthGoodwill = AvgBSSG.GrowthGoodwill / float64(len(BSSG))
	AvgBSSG.GrowthIntangibleAssets = AvgBSSG.GrowthIntangibleAssets / float64(len(BSSG))
	AvgBSSG.GrowthGoodwillAndIntangibleAssets = AvgBSSG.GrowthGoodwillAndIntangibleAssets / float64(len(BSSG))
	AvgBSSG.GrowthLongTermInvestments = AvgBSSG.GrowthLongTermInvestments / float64(len(BSSG))
	AvgBSSG.GrowthTaxAssets = AvgBSSG.GrowthTaxAssets / float64(len(BSSG))
	AvgBSSG.GrowthOtherNonCurrentAssets = AvgBSSG.GrowthOtherNonCurrentAssets / float64(len(BSSG))
	AvgBSSG.GrowthTotalNonCurrentAssets = AvgBSSG.GrowthTotalNonCurrentAssets / float64(len(BSSG))
	AvgBSSG.GrowthOtherAssets = AvgBSSG.GrowthOtherAssets / float64(len(BSSG))
	AvgBSSG.GrowthTotalAssets = AvgBSSG.GrowthTotalAssets / float64(len(BSSG))
	AvgBSSG.GrowthAccountPayables = AvgBSSG.GrowthAccountPayables / float64(len(BSSG))
	AvgBSSG.GrowthShortTermDebt = AvgBSSG.GrowthShortTermDebt / float64(len(BSSG))
	AvgBSSG.GrowthTaxPayables = AvgBSSG.GrowthTaxPayables / float64(len(BSSG))
	AvgBSSG.GrowthDeferredRevenue = AvgBSSG.GrowthDeferredRevenue / float64(len(BSSG))
	AvgBSSG.GrowthOtherCurrentLiabilities = AvgBSSG.GrowthOtherCurrentLiabilities / float64(len(BSSG))
	AvgBSSG.GrowthTotalCurrentLiabilities = AvgBSSG.GrowthTotalCurrentLiabilities / float64(len(BSSG))
	AvgBSSG.GrowthLongTermDebt = AvgBSSG.GrowthLongTermDebt / float64(len(BSSG))
	AvgBSSG.GrowthDeferredRevenueNonCurrent = AvgBSSG.GrowthDeferredRevenueNonCurrent / float64(len(BSSG))
	AvgBSSG.GrowthDeferrredTaxLiabilitiesNonCurrent = AvgBSSG.GrowthDeferrredTaxLiabilitiesNonCurrent / float64(len(BSSG))
	AvgBSSG.GrowthOtherNonCurrentLiabilities = AvgBSSG.GrowthOtherNonCurrentLiabilities / float64(len(BSSG))
	AvgBSSG.GrowthTotalNonCurrentLiabilities = AvgBSSG.GrowthTotalNonCurrentLiabilities / float64(len(BSSG))
	AvgBSSG.GrowthOtherLiabilities = AvgBSSG.GrowthOtherLiabilities / float64(len(BSSG))
	AvgBSSG.GrowthTotalLiabilities = AvgBSSG.GrowthTotalLiabilities / float64(len(BSSG))
	AvgBSSG.GrowthCommonStock = AvgBSSG.GrowthCommonStock / float64(len(BSSG))
	AvgBSSG.GrowthRetainedEarnings = AvgBSSG.GrowthRetainedEarnings / float64(len(BSSG))
	AvgBSSG.GrowthAccumulatedOtherComprehensiveIncomeLoss = AvgBSSG.GrowthAccumulatedOtherComprehensiveIncomeLoss / float64(len(BSSG))
	AvgBSSG.GrowthOthertotalStockholdersEquity = AvgBSSG.GrowthOthertotalStockholdersEquity / float64(len(BSSG))
	AvgBSSG.GrowthTotalStockholdersEquity = AvgBSSG.GrowthTotalStockholdersEquity / float64(len(BSSG))
	AvgBSSG.GrowthTotalLiabilitiesAndStockholdersEquity = AvgBSSG.GrowthTotalLiabilitiesAndStockholdersEquity / float64(len(BSSG))
	AvgBSSG.GrowthTotalInvestments = AvgBSSG.GrowthTotalInvestments / float64(len(BSSG))
	AvgBSSG.GrowthTotalDebt = AvgBSSG.GrowthTotalDebt / float64(len(BSSG))
	AvgBSSG.GrowthNetDebt = AvgBSSG.GrowthNetDebt / float64(len(BSSG))

	return AvgBSSG
}

func GetAverageFinancialStatementGrowth(FSG []objects.FinancialStatementsGrowth) objects.FinancialStatementsGrowth {
	AvgFSG := objects.FinancialStatementsGrowth{}

	for _, Growth := range FSG {
		AvgFSG.RevenueGrowth += Growth.RevenueGrowth
		AvgFSG.GrossProfitGrowth += Growth.GrossProfitGrowth
		AvgFSG.Ebitgrowth += Growth.Ebitgrowth
		AvgFSG.OperatingIncomeGrowth += Growth.OperatingIncomeGrowth
		AvgFSG.NetIncomeGrowth += Growth.NetIncomeGrowth
		AvgFSG.Epsgrowth += Growth.Epsgrowth
		AvgFSG.EpsdilutedGrowth += Growth.EpsdilutedGrowth
		AvgFSG.WeightedAverageSharesGrowth += Growth.WeightedAverageSharesGrowth
		AvgFSG.WeightedAverageSharesDilutedGrowth += Growth.WeightedAverageSharesDilutedGrowth
		AvgFSG.DividendsperShareGrowth += Growth.DividendsperShareGrowth
		AvgFSG.OperatingCashFlowGrowth += Growth.OperatingCashFlowGrowth
		AvgFSG.FreeCashFlowGrowth += Growth.FreeCashFlowGrowth
		AvgFSG.TenYRevenueGrowthPerShare += Growth.TenYRevenueGrowthPerShare
		AvgFSG.FiveYRevenueGrowthPerShare += Growth.FiveYRevenueGrowthPerShare
		AvgFSG.ThreeYRevenueGrowthPerShare += Growth.ThreeYRevenueGrowthPerShare
		AvgFSG.TenYOperatingCFGrowthPerShare += Growth.TenYOperatingCFGrowthPerShare
		AvgFSG.FiveYOperatingCFGrowthPerShare += Growth.FiveYOperatingCFGrowthPerShare
		AvgFSG.ThreeYOperatingCFGrowthPerShare += Growth.ThreeYOperatingCFGrowthPerShare
		AvgFSG.TenYNetIncomeGrowthPerShare += Growth.TenYNetIncomeGrowthPerShare
		AvgFSG.FiveYNetIncomeGrowthPerShare += Growth.FiveYNetIncomeGrowthPerShare
		AvgFSG.ThreeYNetIncomeGrowthPerShare += Growth.ThreeYNetIncomeGrowthPerShare
		AvgFSG.TenYShareholdersEquityGrowthPerShare += Growth.TenYShareholdersEquityGrowthPerShare
		AvgFSG.FiveYShareholdersEquityGrowthPerShare += Growth.FiveYShareholdersEquityGrowthPerShare
		AvgFSG.ThreeYShareholdersEquityGrowthPerShare += Growth.ThreeYShareholdersEquityGrowthPerShare
		AvgFSG.TenYDividendperShareGrowthPerShare += Growth.TenYDividendperShareGrowthPerShare
		AvgFSG.FiveYDividendperShareGrowthPerShare += Growth.FiveYDividendperShareGrowthPerShare
		AvgFSG.ThreeYDividendperShareGrowthPerShare += Growth.ThreeYDividendperShareGrowthPerShare
		AvgFSG.ReceivablesGrowth += Growth.ReceivablesGrowth
		AvgFSG.InventoryGrowth += Growth.InventoryGrowth
		AvgFSG.AssetGrowth += Growth.AssetGrowth
		AvgFSG.BookValueperShareGrowth += Growth.BookValueperShareGrowth
		AvgFSG.DebtGrowth += Growth.DebtGrowth
		AvgFSG.RdexpenseGrowth += Growth.RdexpenseGrowth
		AvgFSG.SgaexpensesGrowth += Growth.SgaexpensesGrowth
	}

	AvgFSG.RevenueGrowth = AvgFSG.RevenueGrowth / float64(len(FSG))
	AvgFSG.GrossProfitGrowth = AvgFSG.GrossProfitGrowth / float64(len(FSG))
	AvgFSG.Ebitgrowth = AvgFSG.Ebitgrowth / float64(len(FSG))
	AvgFSG.OperatingIncomeGrowth = AvgFSG.OperatingIncomeGrowth / float64(len(FSG))
	AvgFSG.NetIncomeGrowth = AvgFSG.NetIncomeGrowth / float64(len(FSG))
	AvgFSG.Epsgrowth = AvgFSG.Epsgrowth / float64(len(FSG))
	AvgFSG.EpsdilutedGrowth = AvgFSG.EpsdilutedGrowth / float64(len(FSG))
	AvgFSG.WeightedAverageSharesGrowth = AvgFSG.WeightedAverageSharesGrowth / float64(len(FSG))
	AvgFSG.WeightedAverageSharesDilutedGrowth = AvgFSG.WeightedAverageSharesDilutedGrowth / float64(len(FSG))
	AvgFSG.DividendsperShareGrowth = AvgFSG.DividendsperShareGrowth / float64(len(FSG))
	AvgFSG.OperatingCashFlowGrowth = AvgFSG.OperatingCashFlowGrowth / float64(len(FSG))
	AvgFSG.FreeCashFlowGrowth = AvgFSG.FreeCashFlowGrowth / float64(len(FSG))
	AvgFSG.TenYRevenueGrowthPerShare = AvgFSG.TenYRevenueGrowthPerShare / float64(len(FSG))
	AvgFSG.FiveYRevenueGrowthPerShare = AvgFSG.FiveYRevenueGrowthPerShare / float64(len(FSG))
	AvgFSG.ThreeYRevenueGrowthPerShare = AvgFSG.ThreeYRevenueGrowthPerShare / float64(len(FSG))
	AvgFSG.TenYOperatingCFGrowthPerShare = AvgFSG.TenYOperatingCFGrowthPerShare / float64(len(FSG))
	AvgFSG.FiveYOperatingCFGrowthPerShare = AvgFSG.FiveYOperatingCFGrowthPerShare / float64(len(FSG))
	AvgFSG.ThreeYOperatingCFGrowthPerShare = AvgFSG.ThreeYOperatingCFGrowthPerShare / float64(len(FSG))
	AvgFSG.TenYNetIncomeGrowthPerShare = AvgFSG.TenYNetIncomeGrowthPerShare / float64(len(FSG))
	AvgFSG.FiveYNetIncomeGrowthPerShare = AvgFSG.FiveYNetIncomeGrowthPerShare / float64(len(FSG))
	AvgFSG.ThreeYNetIncomeGrowthPerShare = AvgFSG.ThreeYNetIncomeGrowthPerShare / float64(len(FSG))
	AvgFSG.TenYShareholdersEquityGrowthPerShare = AvgFSG.TenYShareholdersEquityGrowthPerShare / float64(len(FSG))
	AvgFSG.FiveYShareholdersEquityGrowthPerShare = AvgFSG.FiveYShareholdersEquityGrowthPerShare / float64(len(FSG))
	AvgFSG.ThreeYShareholdersEquityGrowthPerShare = AvgFSG.ThreeYShareholdersEquityGrowthPerShare / float64(len(FSG))
	AvgFSG.TenYDividendperShareGrowthPerShare = AvgFSG.TenYDividendperShareGrowthPerShare / float64(len(FSG))
	AvgFSG.FiveYDividendperShareGrowthPerShare = AvgFSG.FiveYDividendperShareGrowthPerShare / float64(len(FSG))
	AvgFSG.ThreeYDividendperShareGrowthPerShare = AvgFSG.ThreeYDividendperShareGrowthPerShare / float64(len(FSG))
	AvgFSG.ReceivablesGrowth = AvgFSG.ReceivablesGrowth / float64(len(FSG))
	AvgFSG.InventoryGrowth = AvgFSG.InventoryGrowth / float64(len(FSG))
	AvgFSG.AssetGrowth = AvgFSG.AssetGrowth / float64(len(FSG))
	AvgFSG.BookValueperShareGrowth = AvgFSG.BookValueperShareGrowth / float64(len(FSG))
	AvgFSG.DebtGrowth = AvgFSG.DebtGrowth / float64(len(FSG))
	AvgFSG.RdexpenseGrowth = AvgFSG.RdexpenseGrowth / float64(len(FSG))
	AvgFSG.SgaexpensesGrowth = AvgFSG.SgaexpensesGrowth / float64(len(FSG))

	return AvgFSG
}

func GetAverageValuationScore(VS []ValuationScore) ValuationScore {
	AvgVS := ValuationScore{}

	for _, Score := range VS {
		AvgVS.AltmanZScore += Score.AltmanZScore
		AvgVS.PiotroskiScore += Score.PiotroskiScore
	}

	AvgVS.AltmanZScore = AvgVS.AltmanZScore / float64(len(VS))
	AvgVS.PiotroskiScore = AvgVS.PiotroskiScore / float64(len(VS))

	return AvgVS
}
