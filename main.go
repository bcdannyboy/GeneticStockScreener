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

	fmt.Printf("Initiating Genetic Algorithm with mutation rate: %f\n", mutRate)
	ga := genetic.NewGA(mutRate, 1000, 10000, 250, 0.5, 250, FMPAPIClient)

	randAmt := 500
	randomX := make([]string, randAmt)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < randAmt; i++ {
		randomX[i] = tickers.SP500Tickers[rand.Intn(len(tickers.SP500Tickers))]
	}
	ga.RunGeneticAlgorithm(randomX)

	// ga.RunGeneticAlgorithm(tickers.SP500Tickers)
}
