# GeneticStockScreener
Genetic algorithm for screening equities based on fundamental metrics, ratios, and scores.

## Metrics Used

All of the below metrics, as well as their averages and standard deviations across periods (e.g. 5 years, 10 years, etc.), are used as inputs to the genetic algorithm.

### Key Metrics & Key Metrics TTM
Documentation: 
- [Key Metrics](https://site.financialmodelingprep.com/developer/docs#key-metrics-statement-analysis)
- [Key Metrics TTM](https://site.financialmodelingprep.com/developer/docs#key-metrics-ttm-statement-analysis)

Metrics (not including TTM variants, which are the same but trailing twelve months):
- `revenuePerShare`
- `netIncomePerShare`
- `operatingCashFlowPerShare`
- `freeCashFlowPerShare`
- `cashPerShare`
- `bookValuePerShare`
- `tangibleBookValuePerShare`
- `shareholdersEquityPerShare`
- `interestDebtPerShare`
- `marketCap`
- `enterpriseValue`
- `peRatio`
- `priceToSalesRatio`
- `pocfratio`
- `pfcfRatio`
- `pbRatio`
- `ptbRatio`
- `evToSales`
- `enterpriseValueOverEBITDA`
- `evToOperatingCashFlow`
- `evToFreeCashFlow`
- `earningsYield`
- `freeCashFlowYield`
- `debtToEquity`
- `debtToAssets`
- `netDebtToEBITDA`
- `currentRatio`
- `interestCoverage`
- `incomeQuality`
- `dividendYield`
- `payoutRatio`
- `salesGeneralAndAdministrativeToRevenue`
- `researchAndDdevelopementToRevenue`
- `intangiblesToTotalAssets`
- `capexToOperatingCashFlow`
- `capexToRevenue`
- `capexToDepreciation`
- `stockBasedCompensationToRevenue`
- `grahamNumber`
- `roic`
- `returnOnTangibleAssets`
- `grahamNetNet`
- `workingCapital`
- `tangibleAssetValue`
- `netCurrentAssetValue`
- `investedCapital`
- `averageReceivables`
- `averagePayables`
- `averageInventory`
- `daysSalesOutstanding`
- `daysPayablesOutstanding`
- `daysOfInventoryOnHand`
- `receivablesTurnover`
- `payablesTurnover`
- `inventoryTurnover`
- `roe`
- `capexPerShare`

### Ratios
Documentation:
- [Ratios](https://site.financialmodelingprep.com/developer/docs#ratios-statement-analysis)
- [Ratios TTM](https://site.financialmodelingprep.com/developer/docs#ratios-ttm-statement-analysis)

Ratios (not including TTM variants, which are the same but trailing twelve months):
- `currentRatio`
- `quickRatio`
- `cashRatio`
- `daysOfSalesOutstanding`
- `daysOfInventoryOutstanding`
- `operatingCycle`
- `daysOfPayablesOutstanding`
- `cashConversionCycle`
- `grossProfitMargin`
- `operatingProfitMargin`
- `pretaxProfitMargin`
- `netProfitMargin`
- `effectiveTaxRate`
- `returnOnAssets`
- `returnOnEquity`
- `returnOnCapitalEmployed`
- `netIncomePerEBT`
- `ebtPerEbit`
- `ebitPerRevenue`
- `debtRatio`
- `debtEquityRatio`
- `longTermDebtToCapitalization`
- `totalDebtToCapitalization`
- `interestCoverage`
- `cashFlowToDebtRatio`
- `companyEquityMultiplier`
- `receivablesTurnover`
- `payablesTurnover`
- `inventoryTurnover`
- `fixedAssetTurnover`
- `assetTurnover`
- `operatingCashFlowPerShare`
- `freeCashFlowPerShare`
- `cashPerShare`
- `payoutRatio`
- `operatingCashFlowSalesRatio`
- `freeCashFlowOperatingCashFlowRatio`
- `cashFlowCoverageRatios`
- `shortTermCoverageRatios`
- `capitalExpenditureCoverageRatio`
- `dividendPaidAndCapexCoverageRatio`
- `dividendPayoutRatio`
- `priceBookValueRatio`
- `priceToBookRatio`
- `priceToSalesRatio`
- `priceEarningsRatio`
- `priceToFreeCashFlowsRatio`
- `priceToOperatingCashFlowsRatio`
- `priceCashFlowRatio`
- `priceEarningsToGrowthRatio`
- `priceSalesRatio`
- `dividendYield`
- `enterpriseValueMultiple`
- `priceFairValue`

### Cashflow Growth
Documentation:
- [Cashflow Growth](https://site.financialmodelingprep.com/developer/docs#cashflow-growth-statement-analysis)

Metrics:

- `growthNetIncome`
- `growthDepreciationAndAmortization`
- `growthStockBasedCompensation`
- `growthChangeInWorkingCapital`
- `growthAccountsReceivables`
- `growthInventory`
- `growthAccountsPayables`
- `growthOtherWorkingCapital`
- `growthOtherNonCashItems`
- `growthNetCashProvidedByOperatingActivities`
- `growthInvestmentsInPropertyPlantAndEquipment`
- `growthAcquisitionsNet`
- `growthPurchasesOfInvestments`
- `growthSalesMaturitiesOfInvestments`
- `growthNetCashUsedForInvestingActivities`
- `growthDebtRepayment`
- `growthCommonStockIssued`
- `growthCommonStockRepurchased`
- `growthDeferredIncomeTax`
- `growthDividendsPaid`
- `growthNetCashUsedProvidedByFinancingActivities`
- `growthEffectOfForexChangesOnCash`
- `growthNetChangeInCash`
- `growthCashAtEndOfPeriod`
- `growthCashAtBeginningOfPeriod`
- `growthOperatingCashFlow`
- `growthCapitalExpenditure`
- `growthFreeCashFlow`
- `growthOtherInvestingActivites`
- `growthOtherFinancingActivites`

### Income Growth
Documentation:
- [Income Growth](https://site.financialmodelingprep.com/developer/docs#income-growth-statement-analysis)

Metrics:

- `growthRevenue`
- `growthCostOfRevenue`
- `growthGrossProfit`
- `growthGrossProfitRatio`
- `growthResearchAndDevelopmentExpenses`
- `growthGeneralAndAdministrativeExpenses`
- `growthSellingGeneralAndAdministrativeExpenses`
- `growthSellingAndMarketingExpenses`
- `growthOtherExpenses`
- `growthOperatingExpenses`
- `growthCostAndExpenses`
- `growthInterestExpense`
- `growthInterestIncome`
- `growthDepreciationAndAmortization`
- `growthEbitda`
- `growthEbitdaRatio`
- `growthOperatingIncome`
- `growthOperatingIncomeRatio`
- `growthTotalOtherIncomeExpensesNet`
- `growthIncomeBeforeTax`
- `growthIncomeBeforeTaxRatio`
- `growthIncomeTaxExpense`
- `growthNetIncome`
- `growthNetIncomeRatio`
- `growthEps`
- `growthEpsdiluted`
- `growthWeightedAverageShsOut`
- `growthWeightedAverageShsOutDil`

### Balance Sheet Growth
Documentation:
- [Balance Sheet Growth](https://site.financialmodelingprep.com/developer/docs#balance-sheet-growth-statement-analysis)

Metrics:

- `growthCashAndCashEquivalents`
- `growthShortTermInvestments`
- `growthCashAndShortTermInvestments`
- `growthNetReceivables`
- `growthInventory`
- `growthPreferredStock`
- `growthOtherCurrentAssets`
- `growthTotalCurrentAssets`
- `growthPropertyPlantEquipmentNet`
- `growthGoodwill`
- `growthIntangibleAssets`
- `growthGoodwillAndIntangibleAssets`
- `growthLongTermInvestments`
- `growthTaxAssets`
- `growthOtherNonCurrentAssets`
- `growthTotalNonCurrentAssets`
- `growthOtherAssets`
- `growthTotalAssets`
- `growthAccountPayables`
- `growthShortTermDebt`
- `growthTaxPayables`
- `growthCapitalLeaseObligations`
- `growthDeferredRevenue`
- `growthOtherCurrentLiabilities`
- `growthTotalCurrentLiabilities`
- `growthLongTermDebt`
- `growthDeferredRevenueNonCurrent`
- `growthDeferredTaxLiabilitiesNonCurrent`
- `growthOtherNonCurrentLiabilities`
- `growthTotalNonCurrentLiabilities`
- `growthOtherLiabilities`
- `growthTotalLiabilities`
- `growthCommonStock`
- `growthRetainedEarnings`
- `growthAccumulatedOtherComprehensiveIncomeLoss`
- `growthOthertotalStockholdersEquity`
- `growthTotalStockholdersEquity`
- `growthTotalLiabilitiesAndStockholdersEquity`
- `growthTotalInvestments`
- `growthTotalDebt`
- `growthNetDebt`

### Financial Growth
Documentation:
- [Financial Growth](https://site.financialmodelingprep.com/developer/docs#financial-growth-statement-analysis)

Metrics:

- `revenueGrowth`
- `grossProfitGrowth`
- `ebitgrowth`
- `operatingIncomeGrowth`
- `netIncomeGrowth`
- `epsgrowth`
- `epsdilutedGrowth`
- `weightedAverageSharesGrowth`
- `weightedAverageSharesDilutedGrowth`
- `dividendsperShareGrowth`
- `operatingCashFlowGrowth`
- `freeCashFlowGrowth`
- `tenYRevenueGrowthPerShare`
- `fiveYRevenueGrowthPerShare`
- `threeYRevenueGrowthPerShare`
- `tenYOperatingCFGrowthPerShare`
- `fiveYOperatingCFGrowthPerShare`
- `threeYOperatingCFGrowthPerShare`
- `tenYNetIncomeGrowthPerShare`
- `fiveYNetIncomeGrowthPerShare`
- `threeYNetIncomeGrowthPerShare`
- `tenYShareholdersEquityGrowthPerShare`
- `fiveYShareholdersEquityGrowthPerShare`
- `threeYShareholdersEquityGrowthPerShare`
- `tenYDividendperShareGrowthPerShare`
- `fiveYDividendperShareGrowthPerShare`
- `threeYDividendperShareGrowthPerShare`
- `receivablesGrowth`
- `inventoryGrowth`
- `assetGrowth`
- `bookValueperShareGrowth`
- `debtGrowth`
- `rdexpenseGrowth`
- `sgaexpensesGrowth`


## .env Requirements
- !`FMPAPIKEY` - API key for Financial Modeling Prep API