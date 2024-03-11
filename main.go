package main

import (
	"fmt"
	"os"

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

	ga := genetic.NewGA(0.10, 0.05, 10000, 10000, FMPAPIClient)

	// randAmt := 10
	// randomX := make([]string, randAmt)
	// for i := 0; i < randAmt; i++ {
	// 	randomX[i] = tickers.SP500Tickers[rand.Intn(len(tickers.SP500Tickers))]
	// }
	// ga.RunGeneticAlgorithm(randomX)

	ga.RunGeneticAlgorithm(tickers.SP500Tickers)
}
