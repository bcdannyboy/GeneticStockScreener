package genetic

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/bcdannyboy/GeneticStockScreener/src/FMP"
	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

type GA struct {
	MutationRate       float64
	PopulationSize     int
	Generations        int
	APIClient          *FMP.FMPAPI
	TickerFundamentals map[string]*Individual
	BestPortfolio      []string
	mtx                sync.Mutex
}

func (ga *GA) PreFetchFundamentals(TickerPopulation []string) error {
	ga.TickerFundamentals = make(map[string]*Individual)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, runtime.NumCPU()) // Limit goroutines based on CPU count
	perMinuteLimit := 10

	rateLimiter := time.NewTicker(time.Minute / time.Duration(perMinuteLimit))
	defer rateLimiter.Stop()

	fetch := func(ticker string) {
		defer wg.Done()
		<-semaphore                                // Wait for an available slot
		defer func() { semaphore <- struct{}{} }() // Release the slot when done

		individual, err := ga.NewRandomIndividual(ticker)
		if err != nil || individual == nil { // Check for errors and nil individual
			fmt.Printf("Error fetching fundamentals for ticker %s: %v\n", ticker, err)
			return
		}

		ga.mtx.Lock()
		ga.TickerFundamentals[ticker] = individual
		ga.mtx.Unlock()
		fmt.Printf("Fetched fundamentals for %d/%d: %s\n", len(ga.TickerFundamentals), len(TickerPopulation), ticker)
	}

	// Initialize semaphore with the number of concurrent goroutines we want to allow
	for i := 0; i < cap(semaphore); i++ {
		semaphore <- struct{}{}
	}

	for _, ticker := range TickerPopulation {
		wg.Add(1)
		go fetch(ticker)

		// Rate limit block
		<-rateLimiter.C
	}

	wg.Wait()
	fmt.Println("Fundamentals pre-fetched for all tickers.")
	return nil
}

func NewGA(mutationRate float64, populationSize, generations int, APIClient *FMP.FMPAPI) *GA {
	return &GA{
		MutationRate:   mutationRate,
		PopulationSize: populationSize,
		Generations:    generations,
		APIClient:      APIClient,
	}
}

func (ga *GA) NewRandomIndividual(symbol string) (*Individual, error) {
	PriceChange, err := ga.APIClient.GetPriceChange(symbol)
	if err != nil {
		return nil, fmt.Errorf("error getting price change for %s: %v", symbol, err)
	}

	PCScore := genetic_weight.GetPriceChangeScore(PriceChange)
	Fundamentals, err := ga.APIClient.GetValuationInfo(symbol, "quarterly")
	if err != nil {
		return nil, fmt.Errorf("error getting company valuation info for %s: %v", symbol, err)
	}

	w := genetic_weight.InitializeRandomWeight()

	return &Individual{
		Symbol:           symbol,
		PriceChangeScore: PCScore,
		Fundamentals:     Fundamentals,
		Weight:           &w,
	}, nil
}

func topPerformers(population []*Individual, count int) []*Individual {
	// First, sort the population by composite score in descending order.
	sort.Slice(population, func(i, j int) bool {
		return population[i].FundamentalScore > population[j].FundamentalScore
	})

	// Select top performers with unique tickers.
	uniquePerformers := make([]*Individual, 0, count)
	seenTickers := make(map[string]struct{})

	for _, individual := range population {
		if _, seen := seenTickers[individual.Symbol]; !seen && len(uniquePerformers) < count {
			uniquePerformers = append(uniquePerformers, individual)
			seenTickers[individual.Symbol] = struct{}{}
		}
		if len(uniquePerformers) >= count {
			break
		}
	}

	return uniquePerformers
}

func (ga *GA) EvaluatePortfolioFitness(portfolio []*Individual) float64 {
	totalScore := 0.0
	for _, individual := range portfolio {
		totalScore += individual.PriceChangeScore
	}
	averageScore := totalScore / float64(len(portfolio))
	return averageScore
}

func (ga *GA) RunGeneticAlgorithm(TickerPopulation []string) {
	rand.Seed(time.Now().UnixNano())

	if err := ga.PreFetchFundamentals(TickerPopulation); err != nil {
		fmt.Println("Error pre-fetching fundamentals:", err)
		return
	}

	var bestPortfolioFitness float64
	var worstPortfolioFitness float64
	var bestWeights *genetic_weight.Weight
	var mutex sync.Mutex

	for generation := 0; generation < ga.Generations; generation++ {
		fmt.Printf("Initializing Generation %d\n", generation)

		// Initialize population slice with the desired capacity.
		population := make([]*Individual, 0, ga.PopulationSize)

		// Fill the population slice.
		for len(population) < ga.PopulationSize {
			tickerIndex := rand.Intn(len(TickerPopulation))
			ticker := TickerPopulation[tickerIndex]
			individual, exists := ga.TickerFundamentals[ticker]

			if !exists || individual == nil {
				fmt.Printf("Got nil after fetching fundamentals for %s\n", ticker)
				continue
			}

			clonedIndividual := &Individual{
				Symbol:           individual.Symbol,
				PriceChangeScore: individual.PriceChangeScore,
				Fundamentals:     individual.Fundamentals,
				Weight:           genetic_weight.CloneWeights(individual.Weight),
			}

			mutex.Lock()
			population = append(population, clonedIndividual)
			mutex.Unlock()
		}

		// Evaluate fitness in parallel
		var wg sync.WaitGroup
		wg.Add(len(population))
		for _, individual := range population {
			go func(individual *Individual) {
				defer wg.Done()
				individual.FundamentalScore = genetic_weight.CompositeWeightScore(individual.Fundamentals, individual.Weight)
			}(individual)
		}
		wg.Wait()

		topPortfolio := topPerformers(population, 10)
		currentFitness := ga.EvaluatePortfolioFitness(topPortfolio)

		mutex.Lock()
		if currentFitness > bestPortfolioFitness {
			bestPortfolioFitness = currentFitness
			bestWeights = genetic_weight.CloneWeights(topPortfolio[0].Weight)
			ga.BestPortfolio = make([]string, len(topPortfolio))
			for i, individual := range topPortfolio {
				ga.BestPortfolio[i] = individual.Symbol
			}
		} else if currentFitness < worstPortfolioFitness {
			worstPortfolioFitness = currentFitness
		}
		mutex.Unlock()

		// Genetic operations (selection, crossover, mutation)
		for i := 0; i < len(population); i++ {
			parent1, parent2 := ga.SelectWeights(population)
			child := ga.CrossoverWeights(&parent1, &parent2)
			ga.MutateWeights(child)

			mutex.Lock()
			population[i].Weight = child
			mutex.Unlock()
		}

		mutRateMin := (0.001 + 0.005) / 2
		mutRateMax := (0.01 + 0.02) / 2
		mutRate := mutRateMin + rand.Float64()*(mutRateMax-mutRateMin)
		mutRate = mutRate + ga.MutationRate // the last generation should have some effect on the mutation rate
		ga.MutationRate = mutRate / 2
		fmt.Printf("Adjusted mutation rate for next generation (%d): %f\n", generation+1, ga.MutationRate)
	}

	fmt.Printf("Optimization Complete. Best Portfolio Fitness: %f, Worst Portfolio Fitness: %f, Difference: %f, Best to Worst %: %f\n", bestPortfolioFitness, worstPortfolioFitness, bestPortfolioFitness-worstPortfolioFitness, (bestPortfolioFitness-worstPortfolioFitness)/worstPortfolioFitness*100)
	fmt.Println("Top 10 Stocks in the Best Portfolio:")
	for _, symbol := range ga.BestPortfolio {
		fmt.Println(symbol)
	}

	// Save best weights to file
	saveBestWeights(bestWeights)
}

func saveBestWeights(bestWeights *genetic_weight.Weight) {
	jBestWeights, err := json.Marshal(bestWeights)
	if err != nil {
		fmt.Printf("Error marshalling best weights: %v\n", err)
		return
	}

	f := "bestweights.json"
	if err := os.WriteFile(f, jBestWeights, 0644); err != nil {
		fmt.Printf("Error writing best weights to file: %v\n", err)
		return
	}

	fmt.Println("Best weights saved to 'bestweights.json'")
}
