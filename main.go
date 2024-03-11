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

	ga := genetic.NewGA(10000, 10000, FMPAPIClient)

	// random25 := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "FB", "TSLA", "NVDA", "PYPL", "ADBE", "INTC", "CSCO", "NFLX", "CMCSA", "PEP", "AVGO", "COST", "TMUS", "TXN", "QCOM", "AMGN", "SBUX", "CHTR", "INTU", "ISRG", "AMD"}

	ga.RunGeneticAlgorithm(tickers.SP500Tickers)

}
