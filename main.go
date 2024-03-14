package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bcdannyboy/GeneticStockScreener/src/FMP"
	"github.com/bcdannyboy/GeneticStockScreener/src/genetic"
	"github.com/bcdannyboy/GeneticStockScreener/src/tickers"
	"github.com/joho/godotenv"
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

	mutRateMin := (0.001 + 0.005) / 2
	mutRateMax := (0.01 + 0.02) / 2
	mutRate := mutRateMin + rand.Float64()*(mutRateMax-mutRateMin)

	tRates, err := FMPAPIClient.APIClient.Economics.TreasuryRates(time.Now().AddDate(0, 0, -1), time.Now())
	if err != nil {
		panic(fmt.Errorf("error getting treasury rates: %v", err))
	}

	tRate := tRates[0].Year10

	fmt.Printf("Initiating Genetic Algorithm with mutation rate: %f, and 10 year treasury rate: %f\n", mutRate, tRate)
	ga := genetic.NewGA(mutRate, 10000, 10000, 250, 0.5, 250, tRate, 50, FMPAPIClient)

	// randAmt := 25
	// randomX := make([]string, randAmt)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < randAmt; i++ {
	// 	randomX[i] = tickers.SP500Tickers[rand.Intn(len(tickers.SP500Tickers))]
	// }
	// ga.RunGeneticAlgorithm(randomX)

	ga.RunGeneticAlgorithm(tickers.SP500Tickers)
}
