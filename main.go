package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/bcdannyboy/GeneticStockScreener/src/FMP"
	"github.com/bcdannyboy/GeneticStockScreener/src/genetic"
	"github.com/bcdannyboy/GeneticStockScreener/src/tickers"
	"github.com/joho/godotenv"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("no .env file found")
	}

	FMPAPIKEY := os.Getenv("FMPAPIKEY")
	if FMPAPIKEY == "" {
		panic("required environment variable FMPAPIKEY not set")
	}

	FMPAPIClient, err := FMP.NewFMPAPI(FMPAPIKEY)
	if err != nil {
		panic(fmt.Errorf("error creating FMP API client: %v", err))
	}

	mutRateMin := 0.001
	mutRateMax := 0.05
	rand.Seed(time.Now().UnixNano())
	mutRate := mutRateMin + rand.Float64()*(mutRateMax-mutRateMin)

	tRates, err := FMPAPIClient.APIClient.Economics.TreasuryRates(time.Now().AddDate(0, 0, -1), time.Now())
	if err != nil {
		panic(fmt.Errorf("error getting treasury rates: %v", err))
	}

	tRate := tRates[0].Year1
	dailyRiskFreeRate := math.Pow((tRate+1), 1.0/365) - 1

	fmt.Printf("Initiating Genetic Algorithm with mutation rate: %f, and 1 year treasury rate: %f which gives a daily risk free rate of: %f\n", mutRate, tRate, dailyRiskFreeRate)

	// randAmt := 100
	// randomX := make([]string, randAmt)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < randAmt; i++ {
	// 	randomX[i] = tickers.SP500Tickers[rand.Intn(len(tickers.SP500Tickers))]
	// }

	TickerFundamentals := make(map[string]*FMP.CompanyValuationInfo)
	TickerCandles := make(map[string]*objects.StockDailyCandleList)

	// fullTicker := randomX
	fullTicker := tickers.SP500Tickers

	for i, ticker := range fullTicker {
		fundamentals, candles, err := FMPAPIClient.GetValuationInfo(ticker, "quarterly")
		if err != nil {
			fmt.Printf("Error getting valuation info for %s: %v\n", ticker, err)
			continue
		}
		fmt.Printf("Got valuation info for %s (%d/%d)\n", ticker, i+1, len(fullTicker))
		TickerFundamentals[ticker] = fundamentals
		TickerCandles[ticker] = candles
	}

	cpuCount := runtime.NumCPU()
	ga := genetic.NewGA(mutRate, 5000, 10000, 10, 0.5, 500, tRate, 1000, FMPAPIClient, TickerFundamentals, TickerCandles, cpuCount*2)

	topW, bestscore, worstscore, ratio := ga.RunGeneticAlgorithm()
	genetic.SaveBestWeights(topW)
	fmt.Printf("Best Portfolio Score: %f\n", bestscore)
	fmt.Printf("Worst Portfolio Score: %f\n", worstscore)
	fmt.Printf("best/worst ratio: %f\n", ratio)

	ind := &genetic.Individual{Weight: topW}
	ind.PortfolioScore = ga.EvaluateIndividual(ind, true)
	fmt.Printf("Best Portfolio Score (re-evaluated): %f\n", ind.PortfolioScore)
}
