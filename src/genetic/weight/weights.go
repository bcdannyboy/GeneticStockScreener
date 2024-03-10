package genetic_weight

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/bcdannyboy/GeneticStockScreener/src/utils"
)

type W_KeyMetrics struct {
	RevenuePerShare                        float64 `json:"revenuePerShare" csv:"revenuePerShare"`
	NetIncomePerShare                      float64 `json:"netIncomePerShare" csv:"netIncomePerShare"`
	OperatingCashFlowPerShare              float64 `json:"operatingCashFlowPerShare" csv:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare                   float64 `json:"freeCashFlowPerShare" csv:"freeCashFlowPerShare"`
	CashPerShare                           float64 `json:"cashPerShare" csv:"cashPerShare"`
	BookValuePerShare                      float64 `json:"bookValuePerShare" csv:"bookValuePerShare"`
	TangibleBookValuePerShare              float64 `json:"tangibleBookValuePerShare" csv:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare             float64 `json:"shareholdersEquityPerShare" csv:"shareholdersEquityPerShare"`
	InterestDebtPerShare                   float64 `json:"interestDebtPerShare" csv:"interestDebtPerShare"`
	MarketCap                              float64 `json:"marketCap" csv:"marketCap"`
	EnterpriseValue                        float64 `json:"enterpriseValue" csv:"enterpriseValue"`
	PeRatio                                float64 `json:"peRatio" csv:"peRatio"`
	PriceToSalesRatio                      float64 `json:"priceToSalesRatio" csv:"priceToSalesRatio"`
	Pocfratio                              float64 `json:"pocfratio" csv:"pocfratio"`
	PfcfRatio                              float64 `json:"pfcfRatio" csv:"pfcfRatio"`
	PbRatio                                float64 `json:"pbRatio" csv:"pbRatio"`
	PtbRatio                               float64 `json:"ptbRatio" csv:"ptbRatio"`
	EvToSales                              float64 `json:"evToSales" csv:"evToSales"`
	EnterpriseValueOverEBITDA              float64 `json:"enterpriseValueOverEBITDA" csv:"enterpriseValueOverEBITDA"`
	EvToOperatingCashFlow                  float64 `json:"evToOperatingCashFlow" csv:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64 `json:"evToFreeCashFlow" csv:"evToFreeCashFlow"`
	EarningsYield                          float64 `json:"earningsYield" csv:"earningsYield"`
	FreeCashFlowYield                      float64 `json:"freeCashFlowYield" csv:"freeCashFlowYield"`
	DebtToEquity                           float64 `json:"debtToEquity" csv:"debtToEquity"`
	DebtToAssets                           float64 `json:"debtToAssets" csv:"debtToAssets"`
	NetDebtToEBITDA                        float64 `json:"netDebtToEBITDA" csv:"netDebtToEBITDA"`
	CurrentRatio                           float64 `json:"currentRatio" csv:"currentRatio"`
	InterestCoverage                       float64 `json:"interestCoverage" csv:"interestCoverage"`
	IncomeQuality                          float64 `json:"incomeQuality" csv:"incomeQuality"`
	DividendYield                          float64 `json:"dividendYield" csv:"dividendYield"`
	PayoutRatio                            float64 `json:"payoutRatio" csv:"payoutRatio"`
	SalesGeneralAndAdministrativeToRevenue float64 `json:"salesGeneralAndAdministrativeToRevenue" csv:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDdevelopementToRevenue      float64 `json:"researchAndDdevelopementToRevenue" csv:"researchAndDdevelopementToRevenue"`
	IntangiblesToTotalAssets               float64 `json:"intangiblesToTotalAssets" csv:"intangiblesToTotalAssets"`
	CapexToOperatingCashFlow               float64 `json:"capexToOperatingCashFlow" csv:"capexToOperatingCashFlow"`
	CapexToRevenue                         float64 `json:"capexToRevenue" csv:"capexToRevenue"`
	CapexToDepreciation                    float64 `json:"capexToDepreciation" csv:"capexToDepreciation"`
	StockBasedCompensationToRevenue        float64 `json:"stockBasedCompensationToRevenue" csv:"stockBasedCompensationToRevenue"`
	GrahamNumber                           float64 `json:"grahamNumber" csv:"grahamNumber"`
	Roic                                   float64 `json:"roic" csv:"roic"`
	ReturnOnTangibleAssets                 float64 `json:"returnOnTangibleAssets" csv:"returnOnTangibleAssets"`
	GrahamNetNet                           float64 `json:"grahamNetNet" csv:"grahamNetNet"`
	WorkingCapital                         float64 `json:"workingCapital" csv:"workingCapital"`
	TangibleAssetValue                     float64 `json:"tangibleAssetValue" csv:"tangibleAssetValue"`
	NetCurrentAssetValue                   float64 `json:"netCurrentAssetValue" csv:"netCurrentAssetValue"`
	InvestedCapital                        float64 `json:"investedCapital" csv:"investedCapital"`
	AverageReceivables                     float64 `json:"averageReceivables" csv:"averageReceivables"`
	AveragePayables                        float64 `json:"averagePayables" csv:"averagePayables"`
	AverageInventory                       float64 `json:"averageInventory" csv:"averageInventory"`
	DaysSalesOutstanding                   float64 `json:"daysSalesOutstanding" csv:"daysSalesOutstanding"`
	DaysPayablesOutstanding                float64 `json:"daysPayablesOutstanding" csv:"daysPayablesOutstanding"`
	DaysOfInventoryOnHand                  float64 `json:"daysOfInventoryOnHand" csv:"daysOfInventoryOnHand"`
	ReceivablesTurnover                    float64 `json:"receivablesTurnover" csv:"receivablesTurnover"`
	PayablesTurnover                       float64 `json:"payablesTurnover" csv:"payablesTurnover"`
	InventoryTurnover                      float64 `json:"inventoryTurnover" csv:"inventoryTurnover"`
	Roe                                    float64 `json:"roe" csv:"roe"`
	CapexPerShare                          float64 `json:"capexPerShare" csv:"capexPerShare"`
}

type W_KeyMetricsTTM struct {
	RevenuePerShareTTM                        float64 `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM                      float64 `json:"netIncomePerShareTTM"`
	OperatingCashFlowPerShareTTM              float64 `json:"operatingCashFlowPerShareTTM"`
	FreeCashFlowPerShareTTM                   float64 `json:"freeCashFlowPerShareTTM"`
	CashPerShareTTM                           float64 `json:"cashPerShareTTM"`
	BookValuePerShareTTM                      float64 `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM              float64 `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM             float64 `json:"shareholdersEquityPerShareTTM"`
	InterestDebtPerShareTTM                   float64 `json:"interestDebtPerShareTTM"`
	MarketCapTTM                              float64 `json:"marketCapTTM"`
	EnterpriseValueTTM                        float64 `json:"enterpriseValueTTM"`
	PeRatioTTM                                float64 `json:"peRatioTTM"`
	PriceToSalesRatioTTM                      float64 `json:"priceToSalesRatioTTM"`
	PocfratioTTM                              float64 `json:"pocfratioTTM"`
	PfcfRatioTTM                              float64 `json:"pfcfRatioTTM"`
	PbRatioTTM                                float64 `json:"pbRatioTTM"`
	PtbRatioTTM                               float64 `json:"ptbRatioTTM"`
	EvToSalesTTM                              float64 `json:"evToSalesTTM"`
	EnterpriseValueOverEBITDATTM              float64 `json:"enterpriseValueOverEBITDATTM"`
	EvToOperatingCashFlowTTM                  float64 `json:"evToOperatingCashFlowTTM"`
	EvToFreeCashFlowTTM                       float64 `json:"evToFreeCashFlowTTM"`
	EarningsYieldTTM                          float64 `json:"earningsYieldTTM"`
	FreeCashFlowYieldTTM                      float64 `json:"freeCashFlowYieldTTM"`
	DebtToEquityTTM                           float64 `json:"debtToEquityTTM"`
	DividendPerShareTTM                       float64 `json:"dividendPerShareTTM"`
	DebtToAssetsTTM                           float64 `json:"debtToAssetsTTM"`
	NetDebtToEBITDATTM                        float64 `json:"netDebtToEBITDATTM"`
	CurrentRatioTTM                           float64 `json:"currentRatioTTM"`
	InterestCoverageTTM                       float64 `json:"interestCoverageTTM"`
	IncomeQualityTTM                          float64 `json:"incomeQualityTTM"`
	DividendYieldTTM                          float64 `json:"dividendYieldTTM"`
	DividendYieldPercentageTTM                float64 `json:"dividendYieldPercentageTTM"`
	PayoutRatioTTM                            float64 `json:"payoutRatioTTM"`
	SalesGeneralAndAdministrativeToRevenueTTM float64 `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopementToRevenueTTM       float64 `json:"researchAndDevelopementToRevenueTTM"`
	IntangiblesToTotalAssetsTTM               float64 `json:"intangiblesToTotalAssetsTTM"`
	CapexToOperatingCashFlowTTM               float64 `json:"capexToOperatingCashFlowTTM"`
	CapexToRevenueTTM                         float64 `json:"capexToRevenueTTM"`
	CapexToDepreciationTTM                    float64 `json:"capexToDepreciationTTM"`
	StockBasedCompensationToRevenueTTM        float64 `json:"stockBasedCompensationToRevenueTTM"`
	GrahamNumberTTM                           float64 `json:"grahamNumberTTM"`
	RoicTTM                                   float64 `json:"roicTTM"`
	ReturnOnTangibleAssetsTTM                 float64 `json:"returnOnTangibleAssetsTTM"`
	GrahamNetNetTTM                           float64 `json:"grahamNetNetTTM"`
	WorkingCapitalTTM                         float64 `json:"workingCapitalTTM"`
	TangibleAssetValueTTM                     float64 `json:"tangibleAssetValueTTM"`
	NetCurrentAssetValueTTM                   float64 `json:"netCurrentAssetValueTTM"`
	InvestedCapitalTTM                        float64 `json:"investedCapitalTTM"`
	AverageReceivablesTTM                     float64 `json:"averageReceivablesTTM"`
	AveragePayablesTTM                        float64 `json:"averagePayablesTTM"`
	AverageInventoryTTM                       float64 `json:"averageInventoryTTM"`
	DaysSalesOutstandingTTM                   float64 `json:"daysSalesOutstandingTTM"`
	DaysPayablesOutstandingTTM                float64 `json:"daysPayablesOutstandingTTM"`
	DaysOfInventoryOnHandTTM                  float64 `json:"daysOfInventoryOnHandTTM"`
	ReceivablesTurnoverTTM                    float64 `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                       float64 `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                      float64 `json:"inventoryTurnoverTTM"`
	RoeTTM                                    float64 `json:"roeTTM"`
	CapexPerShareTTM                          float64 `json:"capexPerShareTTM"`
}

type W_FinancialRatios struct {
	CurrentRatio                       float64 `json:"currentRatio" csv:"currentRatio"`
	QuickRatio                         float64 `json:"quickRatio" csv:"quickRatio"`
	CashRatio                          float64 `json:"cashRatio" csv:"cashRatio"`
	DaysOfSalesOutstanding             float64 `json:"daysOfSalesOutstanding" csv:"daysOfSalesOutstanding"`
	DaysOfInventoryOutstanding         float64 `json:"daysOfInventoryOutstanding" csv:"daysOfInventoryOutstanding"`
	OperatingCycle                     float64 `json:"operatingCycle" csv:"operatingCycle"`
	DaysOfPayablesOutstanding          float64 `json:"daysOfPayablesOutstanding" csv:"daysOfPayablesOutstanding"`
	CashConversionCycle                float64 `json:"cashConversionCycle" csv:"cashConversionCycle"`
	GrossProfitMargin                  float64 `json:"grossProfitMargin" csv:"grossProfitMargin"`
	OperatingProfitMargin              float64 `json:"operatingProfitMargin" csv:"operatingProfitMargin"`
	PretaxProfitMargin                 float64 `json:"pretaxProfitMargin" csv:"pretaxProfitMargin"`
	NetProfitMargin                    float64 `json:"netProfitMargin" csv:"netProfitMargin"`
	EffectiveTaxRate                   float64 `json:"effectiveTaxRate" csv:"effectiveTaxRate"`
	ReturnOnAssets                     float64 `json:"returnOnAssets" csv:"returnOnAssets"`
	ReturnOnEquity                     float64 `json:"returnOnEquity" csv:"returnOnEquity"`
	ReturnOnCapitalEmployed            float64 `json:"returnOnCapitalEmployed" csv:"returnOnCapitalEmployed"`
	NetIncomePerEBT                    float64 `json:"netIncomePerEBT" csv:"netIncomePerEBT"`
	EbtPerEbit                         float64 `json:"ebtPerEbit" csv:"ebtPerEbit"`
	EbitPerRevenue                     float64 `json:"ebitPerRevenue" csv:"ebitPerRevenue"`
	DebtRatio                          float64 `json:"debtRatio" csv:"debtRatio"`
	DebtEquityRatio                    float64 `json:"debtEquityRatio" csv:"debtEquityRatio"`
	LongTermDebtToCapitalization       float64 `json:"longTermDebtToCapitalization" csv:"longTermDebtToCapitalization"`
	TotalDebtToCapitalization          float64 `json:"totalDebtToCapitalization" csv:"totalDebtToCapitalization"`
	InterestCoverage                   float64 `json:"interestCoverage" csv:"interestCoverage"`
	CashFlowToDebtRatio                float64 `json:"cashFlowToDebtRatio" csv:"cashFlowToDebtRatio"`
	CompanyEquityMultiplier            float64 `json:"companyEquityMultiplier" csv:"companyEquityMultiplier"`
	ReceivablesTurnover                float64 `json:"receivablesTurnover" csv:"receivablesTurnover"`
	PayablesTurnover                   float64 `json:"payablesTurnover" csv:"payablesTurnover"`
	InventoryTurnover                  float64 `json:"inventoryTurnover" csv:"inventoryTurnover"`
	FixedAssetTurnover                 float64 `json:"fixedAssetTurnover" csv:"fixedAssetTurnover"`
	AssetTurnover                      float64 `json:"assetTurnover" csv:"assetTurnover"`
	OperatingCashFlowPerShare          float64 `json:"operatingCashFlowPerShare" csv:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare               float64 `json:"freeCashFlowPerShare" csv:"freeCashFlowPerShare"`
	CashPerShare                       float64 `json:"cashPerShare" csv:"cashPerShare"`
	PayoutRatio                        float64 `json:"payoutRatio" csv:"payoutRatio"`
	OperatingCashFlowSalesRatio        float64 `json:"operatingCashFlowSalesRatio" csv:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio float64 `json:"freeCashFlowOperatingCashFlowRatio" csv:"freeCashFlowOperatingCashFlowRatio"`
	CashFlowCoverageRatios             float64 `json:"cashFlowCoverageRatios" csv:"cashFlowCoverageRatios"`
	ShortTermCoverageRatios            float64 `json:"shortTermCoverageRatios" csv:"shortTermCoverageRatios"`
	CapitalExpenditureCoverageRatio    float64 `json:"capitalExpenditureCoverageRatio" csv:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio  float64 `json:"dividendPaidAndCapexCoverageRatio" csv:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio                float64 `json:"dividendPayoutRatio" csv:"dividendPayoutRatio"`
	PriceBookValueRatio                float64 `json:"priceBookValueRatio" csv:"priceBookValueRatio"`
	PriceToBookRatio                   float64 `json:"priceToBookRatio" csv:"priceToBookRatio"`
	PriceToSalesRatio                  float64 `json:"priceToSalesRatio" csv:"priceToSalesRatio"`
	PriceEarningsRatio                 float64 `json:"priceEarningsRatio" csv:"priceEarningsRatio"`
	PriceToFreeCashFlowsRatio          float64 `json:"priceToFreeCashFlowsRatio" csv:"priceToFreeCashFlowsRatio"`
	PriceToOperatingCashFlowsRatio     float64 `json:"priceToOperatingCashFlowsRatio" csv:"priceToOperatingCashFlowsRatio"`
	PriceCashFlowRatio                 float64 `json:"priceCashFlowRatio" csv:"priceCashFlowRatio"`
	PriceEarningsToGrowthRatio         float64 `json:"priceEarningsToGrowthRatio" csv:"priceEarningsToGrowthRatio"`
	PriceSalesRatio                    float64 `json:"priceSalesRatio" csv:"priceSalesRatio"`
	DividendYield                      float64 `json:"dividendYield" csv:"dividendYield"`
	EnterpriseValueMultiple            float64 `json:"enterpriseValueMultiple" csv:"enterpriseValueMultiple"`
	PriceFairValue                     float64 `json:"priceFairValue" csv:"priceFairValue"`
}

type W_FinancialRatiosTTM struct {
	DividendYielTTM                       float64 `json:"dividendYielTTM" csv:"dividendYielTTM"`
	DividendYielPercentageTTM             float64 `json:"dividendYielPercentageTTM" csv:"dividendYielPercentageTTM"`
	PeRatioTTM                            float64 `json:"peRatioTTM" csv:"peRatioTTM"`
	PegRatioTTM                           float64 `json:"pegRatioTTM" csv:"pegRatioTTM"`
	PayoutRatioTTM                        float64 `json:"payoutRatioTTM" csv:"payoutRatioTTM"`
	CurrentRatioTTM                       float64 `json:"currentRatioTTM" csv:"currentRatioTTM"`
	QuickRatioTTM                         float64 `json:"quickRatioTTM" csv:"quickRatioTTM"`
	CashRatioTTM                          float64 `json:"cashRatioTTM" csv:"cashRatioTTM"`
	DaysOfSalesOutstandingTTM             float64 `json:"daysOfSalesOutstandingTTM" csv:"daysOfSalesOutstandingTTM"`
	DaysOfInventoryOutstandingTTM         float64 `json:"daysOfInventoryOutstandingTTM" csv:"daysOfInventoryOutstandingTTM"`
	OperatingCycleTTM                     float64 `json:"operatingCycleTTM" csv:"operatingCycleTTM"`
	DaysOfPayablesOutstandingTTM          float64 `json:"daysOfPayablesOutstandingTTM" csv:"daysOfPayablesOutstandingTTM"`
	CashConversionCycleTTM                float64 `json:"cashConversionCycleTTM" csv:"cashConversionCycleTTM"`
	GrossProfitMarginTTM                  float64 `json:"grossProfitMarginTTM" csv:"grossProfitMarginTTM"`
	OperatingProfitMarginTTM              float64 `json:"operatingProfitMarginTTM" csv:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM                 float64 `json:"pretaxProfitMarginTTM" csv:"pretaxProfitMarginTTM"`
	NetProfitMarginTTM                    float64 `json:"netProfitMarginTTM" csv:"netProfitMarginTTM"`
	EffectiveTaxRateTTM                   float64 `json:"effectiveTaxRateTTM" csv:"effectiveTaxRateTTM"`
	ReturnOnAssetsTTM                     float64 `json:"returnOnAssetsTTM" csv:"returnOnAssetsTTM"`
	ReturnOnEquityTTM                     float64 `json:"returnOnEquityTTM" csv:"returnOnEquityTTM"`
	ReturnOnCapitalEmployedTTM            float64 `json:"returnOnCapitalEmployedTTM" csv:"returnOnCapitalEmployedTTM"`
	NetIncomePerEBTTTM                    float64 `json:"netIncomePerEBTTTM" csv:"netIncomePerEBTTTM"`
	EbtPerEbitTTM                         float64 `json:"ebtPerEbitTTM" csv:"ebtPerEbitTTM"`
	EbitPerRevenueTTM                     float64 `json:"ebitPerRevenueTTM" csv:"ebitPerRevenueTTM"`
	DebtRatioTTM                          float64 `json:"debtRatioTTM" csv:"debtRatioTTM"`
	DebtEquityRatioTTM                    float64 `json:"debtEquityRatioTTM" csv:"debtEquityRatioTTM"`
	LongTermDebtToCapitalizationTTM       float64 `json:"longTermDebtToCapitalizationTTM" csv:"longTermDebtToCapitalizationTTM"`
	TotalDebtToCapitalizationTTM          float64 `json:"totalDebtToCapitalizationTTM" csv:"totalDebtToCapitalizationTTM"`
	InterestCoverageTTM                   float64 `json:"interestCoverageTTM" csv:"interestCoverageTTM"`
	CashFlowToDebtRatioTTM                float64 `json:"cashFlowToDebtRatioTTM" csv:"cashFlowToDebtRatioTTM"`
	CompanyEquityMultiplierTTM            float64 `json:"companyEquityMultiplierTTM" csv:"companyEquityMultiplierTTM"`
	ReceivablesTurnoverTTM                float64 `json:"receivablesTurnoverTTM" csv:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                   float64 `json:"payablesTurnoverTTM" csv:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                  float64 `json:"inventoryTurnoverTTM" csv:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM                 float64 `json:"fixedAssetTurnoverTTM" csv:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM                      float64 `json:"assetTurnoverTTM" csv:"assetTurnoverTTM"`
	OperatingCashFlowPerShareTTM          float64 `json:"operatingCashFlowPerShareTTM" csv:"operatingCashFlowPerShareTTM"`
	FreeCashFlowPerShareTTM               float64 `json:"freeCashFlowPerShareTTM" csv:"freeCashFlowPerShareTTM"`
	CashPerShareTTM                       float64 `json:"cashPerShareTTM" csv:"cashPerShareTTM"`
	OperatingCashFlowSalesRatioTTM        float64 `json:"operatingCashFlowSalesRatioTTM" csv:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM float64 `json:"freeCashFlowOperatingCashFlowRatioTTM" csv:"freeCashFlowOperatingCashFlowRatioTTM"`
	CashFlowCoverageRatiosTTM             float64 `json:"cashFlowCoverageRatiosTTM" csv:"cashFlowCoverageRatiosTTM"`
	ShortTermCoverageRatiosTTM            float64 `json:"shortTermCoverageRatiosTTM" csv:"shortTermCoverageRatiosTTM"`
	CapitalExpenditureCoverageRatioTTM    float64 `json:"capitalExpenditureCoverageRatioTTM" csv:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTM  float64 `json:"dividendPaidAndCapexCoverageRatioTTM" csv:"dividendPaidAndCapexCoverageRatioTTM"`
	PriceBookValueRatioTTM                float64 `json:"priceBookValueRatioTTM" csv:"priceBookValueRatioTTM"`
	PriceToBookRatioTTM                   float64 `json:"priceToBookRatioTTM" csv:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                  float64 `json:"priceToSalesRatioTTM" csv:"priceToSalesRatioTTM"`
	PriceEarningsRatioTTM                 float64 `json:"priceEarningsRatioTTM" csv:"priceEarningsRatioTTM"`
	PriceToFreeCashFlowsRatioTTM          float64 `json:"priceToFreeCashFlowsRatioTTM" csv:"priceToFreeCashFlowsRatioTTM"`
	PriceToOperatingCashFlowsRatioTTM     float64 `json:"priceToOperatingCashFlowsRatioTTM" csv:"priceToOperatingCashFlowsRatioTTM"`
	PriceCashFlowRatioTTM                 float64 `json:"priceCashFlowRatioTTM" csv:"priceCashFlowRatioTTM"`
	PriceEarningsToGrowthRatioTTM         float64 `json:"priceEarningsToGrowthRatioTTM" csv:"priceEarningsToGrowthRatioTTM"`
	PriceSalesRatioTTM                    float64 `json:"priceSalesRatioTTM" csv:"priceSalesRatioTTM"`
	DividendYieldTTM                      float64 `json:"dividendYieldTTM" csv:"dividendYieldTTM"`
	EnterpriseValueMultipleTTM            float64 `json:"enterpriseValueMultipleTTM" csv:"enterpriseValueMultipleTTM"`
	PriceFairValueTTM                     float64 `json:"priceFairValueTTM" csv:"priceFairValueTTM"`
	DividendPerShareTTM                   float64 `json:"dividendPerShareTTM" csv:"dividendPerShareTTM"`
}

type W_CashFlowStatementGrowth struct {
	GrowthNetIncome                                float64 `json:"growthNetIncome"`
	GrowthDepreciationAndAmortization              float64 `json:"growthDepreciationAndAmortization"`
	GrowthDeferredIncomeTax                        float64 `json:"growthDeferredIncomeTax"`
	GrowthStockBasedCompensation                   float64 `json:"growthStockBasedCompensation"`
	GrowthChangeInWorkingCapital                   float64 `json:"growthChangeInWorkingCapital"`
	GrowthAccountsReceivables                      float64 `json:"growthAccountsReceivables"`
	GrowthInventory                                float64 `json:"growthInventory"`
	GrowthAccountsPayables                         float64 `json:"growthAccountsPayables"`
	GrowthOtherWorkingCapital                      float64 `json:"growthOtherWorkingCapital"`
	GrowthOtherNonCashItems                        float64 `json:"growthOtherNonCashItems"`
	GrowthNetCashProvidedByOperatingActivites      float64 `json:"growthNetCashProvidedByOperatingActivites"`
	GrowthInvestmentsInPropertyPlantAndEquipment   float64 `json:"growthInvestmentsInPropertyPlantAndEquipment"`
	GrowthAcquisitionsNet                          float64 `json:"growthAcquisitionsNet"`
	GrowthPurchasesOfInvestments                   float64 `json:"growthPurchasesOfInvestments"`
	GrowthSalesMaturitiesOfInvestments             float64 `json:"growthSalesMaturitiesOfInvestments"`
	GrowthOtherInvestingActivites                  float64 `json:"growthOtherInvestingActivites"`
	GrowthNetCashUsedForInvestingActivites         float64 `json:"growthNetCashUsedForInvestingActivites"`
	GrowthDebtRepayment                            float64 `json:"growthDebtRepayment"`
	GrowthCommonStockIssued                        float64 `json:"growthCommonStockIssued"`
	GrowthCommonStockRepurchased                   float64 `json:"growthCommonStockRepurchased"`
	GrowthDividendsPaid                            float64 `json:"growthDividendsPaid"`
	GrowthOtherFinancingActivites                  float64 `json:"growthOtherFinancingActivites"`
	GrowthNetCashUsedProvidedByFinancingActivities float64 `json:"growthNetCashUsedProvidedByFinancingActivities"`
	GrowthEffectOfForexChangesOnCash               float64 `json:"growthEffectOfForexChangesOnCash"`
	GrowthNetChangeInCash                          float64 `json:"growthNetChangeInCash"`
	GrowthCashAtEndOfPeriod                        float64 `json:"growthCashAtEndOfPeriod"`
	GrowthCashAtBeginningOfPeriod                  float64 `json:"growthCashAtBeginningOfPeriod"`
	GrowthOperatingCashFlow                        float64 `json:"growthOperatingCashFlow"`
	GrowthCapitalExpenditure                       float64 `json:"growthCapitalExpenditure"`
	GrowthFreeCashFlow                             float64 `json:"growthFreeCashFlow"`
}

type W_IncomeStatementGrowth struct {
	GrowthRevenue                          float64 `json:"growthRevenue"`
	GrowthCostOfRevenue                    float64 `json:"growthCostOfRevenue"`
	GrowthGrossProfit                      float64 `json:"growthGrossProfit"`
	GrowthGrossProfitRatio                 float64 `json:"growthGrossProfitRatio"`
	GrowthResearchAndDevelopmentExpenses   float64 `json:"growthResearchAndDevelopmentExpenses"`
	GrowthGeneralAndAdministrativeExpenses float64 `json:"growthGeneralAndAdministrativeExpenses"`
	GrowthSellingAndMarketingExpenses      float64 `json:"growthSellingAndMarketingExpenses"`
	GrowthOtherExpenses                    float64 `json:"growthOtherExpenses"`
	GrowthOperatingExpenses                float64 `json:"growthOperatingExpenses"`
	GrowthCostAndExpenses                  float64 `json:"growthCostAndExpenses"`
	GrowthInterestExpense                  float64 `json:"growthInterestExpense"`
	GrowthDepreciationAndAmortization      float64 `json:"growthDepreciationAndAmortization"`
	GrowthEBITDA                           float64 `json:"growthEBITDA"`
	GrowthEBITDARatio                      float64 `json:"growthEBITDARatio"`
	GrowthOperatingIncome                  float64 `json:"growthOperatingIncome"`
	GrowthOperatingIncomeRatio             float64 `json:"growthOperatingIncomeRatio"`
	GrowthTotalOtherIncomeExpensesNet      float64 `json:"growthTotalOtherIncomeExpensesNet"`
	GrowthIncomeBeforeTax                  float64 `json:"growthIncomeBeforeTax"`
	GrowthIncomeBeforeTaxRatio             float64 `json:"growthIncomeBeforeTaxRatio"`
	GrowthIncomeTaxExpense                 float64 `json:"growthIncomeTaxExpense"`
	GrowthNetIncome                        float64 `json:"growthNetIncome"`
	GrowthNetIncomeRatio                   float64 `json:"growthNetIncomeRatio"`
	GrowthEPS                              float64 `json:"growthEPS"`
	GrowthEPSDiluted                       float64 `json:"growthEPSDiluted"`
	GrowthWeightedAverageShsOut            float64 `json:"growthWeightedAverageShsOut"`
	GrowthWeightedAverageShsOutDil         float64 `json:"growthWeightedAverageShsOutDil"`
}

type W_BalanceSheetStatementGrowth struct {
	GrowthCashAndCashEquivalents                  float64 `json:"growthCashAndCashEquivalents"`
	GrowthShortTermInvestments                    float64 `json:"growthShortTermInvestments"`
	GrowthCashAndShortTermInvestments             float64 `json:"growthCashAndShortTermInvestments"`
	GrowthNetReceivables                          float64 `json:"growthNetReceivables"`
	GrowthInventory                               float64 `json:"growthInventory"`
	GrowthOtherCurrentAssets                      float64 `json:"growthOtherCurrentAssets"`
	GrowthTotalCurrentAssets                      float64 `json:"growthTotalCurrentAssets"`
	GrowthPropertyPlantEquipmentNet               float64 `json:"growthPropertyPlantEquipmentNet"`
	GrowthGoodwill                                float64 `json:"growthGoodwill"`
	GrowthIntangibleAssets                        float64 `json:"growthIntangibleAssets"`
	GrowthGoodwillAndIntangibleAssets             float64 `json:"growthGoodwillAndIntangibleAssets"`
	GrowthLongTermInvestments                     float64 `json:"growthLongTermInvestments"`
	GrowthTaxAssets                               float64 `json:"growthTaxAssets"`
	GrowthOtherNonCurrentAssets                   float64 `json:"growthOtherNonCurrentAssets"`
	GrowthTotalNonCurrentAssets                   float64 `json:"growthTotalNonCurrentAssets"`
	GrowthOtherAssets                             float64 `json:"growthOtherAssets"`
	GrowthTotalAssets                             float64 `json:"growthTotalAssets"`
	GrowthAccountPayables                         float64 `json:"growthAccountPayables"`
	GrowthShortTermDebt                           float64 `json:"growthShortTermDebt"`
	GrowthTaxPayables                             float64 `json:"growthTaxPayables"`
	GrowthDeferredRevenue                         float64 `json:"growthDeferredRevenue"`
	GrowthOtherCurrentLiabilities                 float64 `json:"growthOtherCurrentLiabilities"`
	GrowthTotalCurrentLiabilities                 float64 `json:"growthTotalCurrentLiabilities"`
	GrowthLongTermDebt                            float64 `json:"growthLongTermDebt"`
	GrowthDeferredRevenueNonCurrent               float64 `json:"growthDeferredRevenueNonCurrent"`
	GrowthDeferrredTaxLiabilitiesNonCurrent       float64 `json:"growthDeferrredTaxLiabilitiesNonCurrent"`
	GrowthOtherNonCurrentLiabilities              float64 `json:"growthOtherNonCurrentLiabilities"`
	GrowthTotalNonCurrentLiabilities              float64 `json:"growthTotalNonCurrentLiabilities"`
	GrowthOtherLiabilities                        float64 `json:"growthOtherLiabilities"`
	GrowthTotalLiabilities                        float64 `json:"growthTotalLiabilities"`
	GrowthCommonStock                             float64 `json:"growthCommonStock"`
	GrowthRetainedEarnings                        float64 `json:"growthRetainedEarnings"`
	GrowthAccumulatedOtherComprehensiveIncomeLoss float64 `json:"growthAccumulatedOtherComprehensiveIncomeLoss"`
	GrowthOthertotalStockholdersEquity            float64 `json:"growthOthertotalStockholdersEquity"`
	GrowthTotalStockholdersEquity                 float64 `json:"growthTotalStockholdersEquity"`
	GrowthTotalLiabilitiesAndStockholdersEquity   float64 `json:"growthTotalLiabilitiesAndStockholdersEquity"`
	GrowthTotalInvestments                        float64 `json:"growthTotalInvestments"`
	GrowthTotalDebt                               float64 `json:"growthTotalDebt"`
	GrowthNetDebt                                 float64 `json:"growthNetDebt"`
}

type W_FinancialStatementsGrowth struct {
	RevenueGrowth                          float64 `json:"revenueGrowth"`
	GrossProfitGrowth                      float64 `json:"grossProfitGrowth"`
	Ebitgrowth                             float64 `json:"ebitgrowth"`
	OperatingIncomeGrowth                  float64 `json:"operatingIncomeGrowth"`
	NetIncomeGrowth                        float64 `json:"netIncomeGrowth"`
	Epsgrowth                              float64 `json:"epsgrowth"`
	EpsdilutedGrowth                       float64 `json:"epsdilutedGrowth"`
	WeightedAverageSharesGrowth            float64 `json:"weightedAverageSharesGrowth"`
	WeightedAverageSharesDilutedGrowth     float64 `json:"weightedAverageSharesDilutedGrowth"`
	DividendsperShareGrowth                float64 `json:"dividendsperShareGrowth"`
	OperatingCashFlowGrowth                float64 `json:"operatingCashFlowGrowth"`
	FreeCashFlowGrowth                     float64 `json:"freeCashFlowGrowth"`
	TenYRevenueGrowthPerShare              float64 `json:"tenYRevenueGrowthPerShare"`
	FiveYRevenueGrowthPerShare             float64 `json:"fiveYRevenueGrowthPerShare"`
	ThreeYRevenueGrowthPerShare            float64 `json:"threeYRevenueGrowthPerShare"`
	TenYOperatingCFGrowthPerShare          float64 `json:"tenYOperatingCFGrowthPerShare"`
	FiveYOperatingCFGrowthPerShare         float64 `json:"fiveYOperatingCFGrowthPerShare"`
	ThreeYOperatingCFGrowthPerShare        float64 `json:"threeYOperatingCFGrowthPerShare"`
	TenYNetIncomeGrowthPerShare            float64 `json:"tenYNetIncomeGrowthPerShare"`
	FiveYNetIncomeGrowthPerShare           float64 `json:"fiveYNetIncomeGrowthPerShare"`
	ThreeYNetIncomeGrowthPerShare          float64 `json:"threeYNetIncomeGrowthPerShare"`
	TenYShareholdersEquityGrowthPerShare   float64 `json:"tenYShareholdersEquityGrowthPerShare"`
	FiveYShareholdersEquityGrowthPerShare  float64 `json:"fiveYShareholdersEquityGrowthPerShare"`
	ThreeYShareholdersEquityGrowthPerShare float64 `json:"threeYShareholdersEquityGrowthPerShare"`
	TenYDividendperShareGrowthPerShare     float64 `json:"tenYDividendperShareGrowthPerShare"`
	FiveYDividendperShareGrowthPerShare    float64 `json:"fiveYDividendperShareGrowthPerShare"`
	ThreeYDividendperShareGrowthPerShare   float64 `json:"threeYDividendperShareGrowthPerShare"`
	ReceivablesGrowth                      float64 `json:"receivablesGrowth"`
	InventoryGrowth                        float64 `json:"inventoryGrowth"`
	AssetGrowth                            float64 `json:"assetGrowth"`
	BookValueperShareGrowth                float64 `json:"bookValueperShareGrowth"`
	DebtGrowth                             float64 `json:"debtGrowth"`
	RdexpenseGrowth                        float64 `json:"rdexpenseGrowth"`
	SgaexpensesGrowth                      float64 `json:"sgaexpensesGrowth"`
}

type Weight struct {
	KeyMetrics                     W_KeyMetrics                  `json:"keyMetrics"`
	AvgKeyMetrics                  W_KeyMetrics                  `json:"avgKeyMetrics"`
	KeyMetricsTTM                  W_KeyMetricsTTM               `json:"keyMetricsTTM"`
	Ratios                         W_FinancialRatios             `json:"financialRatios"`
	AvgRatios                      W_FinancialRatios             `json:"avgFinancialRatios"`
	RatiosTTM                      W_FinancialRatiosTTM          `json:"financialRatiosTTM"`
	CashFlowStatementGrowth        W_CashFlowStatementGrowth     `json:"cashFlowStatementGrowth"`
	AvgCashFlowStatementGrowth     W_CashFlowStatementGrowth     `json:"avgCashFlowStatementGrowth"`
	IncomeStatementGrowth          W_IncomeStatementGrowth       `json:"incomeStatementGrowth"`
	AvgIncomeStatementGrowth       W_IncomeStatementGrowth       `json:"avgIncomeStatementGrowth"`
	BalanceSheetStatementGrowth    W_BalanceSheetStatementGrowth `json:"balanceSheetStatementGrowth"`
	AvgBalanceSheetStatementGrowth W_BalanceSheetStatementGrowth `json:"avgBalanceSheetStatementGrowth"`
	FinancialStatementsGrowth      W_FinancialStatementsGrowth   `json:"financialStatementsGrowth"`
	AvgFinancialStatementsGrowth   W_FinancialStatementsGrowth   `json:"avgFinancialStatementsGrowth"`
}

func InitializeRandomWeight() Weight {
	rand.Seed(time.Now().UnixNano())

	var weight Weight
	val := reflect.ValueOf(&weight).Elem()

	utils.SetRandomFloats(val)

	return weight
}
