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

	ga := genetic.NewGA(0.10, 0.05, 1000, 1000, FMPAPIClient)

	randAmt := 25
	randomX := make([]string, randAmt)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < randAmt; i++ {
		randomX[i] = tickers.SP500Tickers[rand.Intn(len(tickers.SP500Tickers))]
	}
	ga.RunGeneticAlgorithm(randomX)

	// ga.RunGeneticAlgorithm(tickers.SP500Tickers)
}
