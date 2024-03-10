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
	PopulationSize int
	Generations    int
	APIClient      *FMP.FMPAPI
}

func NewGA(populationSize, generations int, APIClient *FMP.FMPAPI) *GA {
	return &GA{
		PopulationSize: populationSize,
		Generations:    generations,
		APIClient:      APIClient,
	}
}

func (ga *GA) GenerateAllIndividuals(TickerPopulation []string) ([]*Individual, error) {
	fmt.Printf("Generating initial population of %d individuals\n", len(TickerPopulation))
	rand.Seed(time.Now().UnixNano())

	individuals := make([]*Individual, len(TickerPopulation))
	var wg sync.WaitGroup
	// A buffered channel to limit goroutines
	semaphore := make(chan struct{}, runtime.NumCPU())

	// Create a ticker for rate limiting to 10 operations per minute
	ticker := time.NewTicker(time.Minute / 10)
	defer ticker.Stop()

	for i, symbol := range TickerPopulation {
		wg.Add(1)
		go func(i int, symbol string) {
			semaphore <- struct{}{} // Acquire a token
			defer wg.Done()
			defer func() { <-semaphore }() // Release the token

			<-ticker.C // Wait for the next tick to enforce rate limit
			individual, err := ga.NewRandomIndividual(symbol)
			if err != nil {
				fmt.Println("Error creating individual:", err)
				return
			}

			// Initialize unique weights for each individual here, inside the loop
			individualWeights := genetic_weight.InitializeRandomWeight()
			individual.Weight = &individualWeights

			fmt.Printf("Generated individual %d: %s with unique weights\n", i, individual.Symbol)
			individuals[i] = individual
		}(i, symbol)
	}
	wg.Wait()
	return individuals, nil
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

	return &Individual{
		Symbol:           symbol,
		PriceChangeScore: PCScore,
		Fundamentals:     Fundamentals,
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

func (ga *GA) RunGeneticAlgorithm(PopulationSource []string) {
	rand.Seed(time.Now().UnixNano())

	// Initialize a set of weights for the entire population
	initialWeights := genetic_weight.InitializeRandomWeight()

	// Pre-generate all individuals with initial weights
	population, err := ga.GenerateAllIndividuals(PopulationSource)
	if err != nil {
		fmt.Printf("Error generating initial population: %v\n", err)
		return
	}

	var bestPortfolioFitness float64
	var bestWeights *genetic_weight.Weight

	// Main optimization loop
	for generation := 0; generation < ga.Generations; generation++ {
		fmt.Printf("Initializing Generation %d\n", generation)

		// Evaluate composite scores in parallel
		var wg sync.WaitGroup
		wg.Add(len(population))
		semaphore := make(chan struct{}, runtime.NumCPU())

		for _, individual := range population {
			go func(individual *Individual) {
				semaphore <- struct{}{}
				defer wg.Done()
				defer func() { <-semaphore }()

				// Apply the same weights to evaluate each individual
				individual.FundamentalScore = genetic_weight.CompositeWeightScore(individual.Fundamentals, &initialWeights)
			}(individual)
		}
		wg.Wait()

		// Select the top performers with unique tickers
		topPortfolio := topPerformers(population, 10)

		// Evaluate the fitness of the portfolio
		currentFitness := ga.EvaluatePortfolioFitness(topPortfolio)
		fmt.Printf("Generation %d: Portfolio Fitness: %f\n", generation, currentFitness)

		// Check if this generation's portfolio is the best so far
		if currentFitness > bestPortfolioFitness {
			bestPortfolioFitness = currentFitness
			bestWeights = &initialWeights // Consider cloning the weights if they are subject to change
		}

		// Selection, crossover, and mutation of weights for the next generation
		// Assume SelectWeights, CrossoverWeights, and MutateWeights are implemented
		parentWeights1, parentWeights2 := ga.SelectWeights(population)
		child := ga.CrossoverWeights(&parentWeights1, &parentWeights2)
		ga.MutateWeights(child)

		// Prepare for the next generation
		initialWeights = *child
	}

	fmt.Printf("Optimization Complete. Best Portfolio Fitness: %f\n", bestPortfolioFitness)

	jBestWeights, err := json.Marshal(bestWeights)
	if err != nil {
		fmt.Printf("Error marshalling best weights: %v\n", err)
		fmt.Printf("Best Weights:\n%v\n", bestWeights)
		return
	}

	f := "bestweights.json"
	err = os.WriteFile(f, jBestWeights, 0644)
	if err != nil {
		fmt.Printf("Error writing best weights to file: %v\n", err)
		fmt.Printf("Best Weights:\n%v\n", bestWeights)
		return
	}
}
