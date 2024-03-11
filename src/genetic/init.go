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
	PopulationSize     int
	Generations        int
	APIClient          *FMP.FMPAPI
	TickerFundamentals map[string]*Individual
	BestPortfolio      []string
}

func (ga *GA) PreFetchFundamentals(TickerPopulation []string) error {
	ga.TickerFundamentals = make(map[string]*Individual, len(TickerPopulation))
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, runtime.NumCPU())

	// Create a ticker that ticks every 6 seconds.
	rateLimiter := time.NewTicker(6 * time.Second)
	defer rateLimiter.Stop()

	// Use a channel to control the rate of API calls.
	requests := make(chan struct{}, 10)

	// Pre-fill the requests channel to allow up to 10 immediate requests.
	for i := 0; i < 10; i++ {
		requests <- struct{}{}
	}

	for _, ticker := range TickerPopulation {
		wg.Add(1)
		go func(ticker string) {
			defer wg.Done()

			// Wait for the ability to make a request or for the next tick.
			select {
			case <-requests:
				// Allowed to make a request.
			case <-rateLimiter.C:
				// Refill the requests channel every tick (6 seconds) for rate-limiting.
				for i := 0; i < 10; i++ {
					select {
					case requests <- struct{}{}:
					default:
						// Channel is already full.
					}
				}
				<-requests // Now we can proceed after ensuring rate limiting.
			}

			semaphore <- struct{}{}        // Acquire a token to limit concurrency.
			defer func() { <-semaphore }() // Release the token.

			individual, err := ga.NewRandomIndividual(ticker)
			if err != nil {
				fmt.Printf("Error fetching fundamentals for ticker %s: %v\n", ticker, err)
				return
			}
			fmt.Printf("Fetched fundamentals for %s\n", ticker)
			ga.TickerFundamentals[ticker] = individual
		}(ticker)
	}
	wg.Wait()
	return nil
}

func NewGA(populationSize, generations int, APIClient *FMP.FMPAPI) *GA {
	return &GA{
		PopulationSize: populationSize,
		Generations:    generations,
		APIClient:      APIClient,
	}
}

func (ga *GA) GenerateAllIndividuals(TickerPopulation []string) ([]*Individual, error) {
	fmt.Printf("Generating initial population\n")
	rand.Seed(time.Now().UnixNano())

	// Determine the actual size of the initial population
	popSize := ga.PopulationSize
	if len(TickerPopulation) > popSize {
		popSize = len(TickerPopulation)
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, runtime.NumCPU())
	tickerRateLimit := time.NewTicker(time.Minute / 10)
	defer tickerRateLimit.Stop()

	// Creating a slice to hold the generated individuals
	individuals := make([]*Individual, 0, popSize)

	for i := 0; i < ga.PopulationSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire a token
			defer func() { <-semaphore }() // Release the token
			<-tickerRateLimit.C            // Rate limit each API call

			// Select a random ticker from the array
			symbol := TickerPopulation[rand.Intn(len(TickerPopulation))]

			individual, err := ga.NewRandomIndividual(symbol)
			if err != nil {
				fmt.Printf("Error creating individual: %v\n", err)
				return
			}

			individualWeights := genetic_weight.InitializeRandomWeight()
			individual.Weight = &individualWeights
			individuals = append(individuals, individual)

			fmt.Printf("Generated individual %d: %s with unique weights\n", i, individual.Symbol)
		}(i)
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

func (ga *GA) RunGeneticAlgorithm(TickerPopulation []string) {
	rand.Seed(time.Now().UnixNano())

	// Ensure the fundamentals for all tickers are pre-fetched before the GA starts
	if err := ga.PreFetchFundamentals(TickerPopulation); err != nil {
		fmt.Println("Error pre-fetching fundamentals:", err)
		return
	}

	var bestPortfolioFitness float64
	var bestWeights *genetic_weight.Weight

	for generation := 0; generation < ga.Generations; generation++ {
		fmt.Printf("Initializing Generation %d\n", generation)

		// Randomize ticker selection for this generation to ensure diversity
		currentGenerationTickers := make([]string, len(TickerPopulation))
		copy(currentGenerationTickers, TickerPopulation)
		rand.Shuffle(len(currentGenerationTickers), func(i, j int) {
			currentGenerationTickers[i], currentGenerationTickers[j] = currentGenerationTickers[j], currentGenerationTickers[i]
		})

		// Generate the population for this generation based on shuffled tickers
		population := make([]*Individual, ga.PopulationSize)
		for i := 0; i < ga.PopulationSize; i++ {
			ticker := currentGenerationTickers[i%len(currentGenerationTickers)]
			individual := ga.TickerFundamentals[ticker]
			// Clone the individual to avoid modifying the original in TickerFundamentals
			clonedIndividual := *individual
			ciWeight := genetic_weight.InitializeRandomWeight()
			clonedIndividual.Weight = &ciWeight
			population[i] = &clonedIndividual
		}

		// Evaluate the population and calculate fitness
		for _, individual := range population {
			individual.FundamentalScore = genetic_weight.CompositeWeightScore(individual.Fundamentals, individual.Weight)
		}

		// Select the top performers
		topPortfolio := topPerformers(population, 10)

		// Evaluate the fitness of the top portfolio
		currentFitness := ga.EvaluatePortfolioFitness(topPortfolio)
		fmt.Printf("Generation %d [%d]: Portfolio Fitness: %f\n", generation, ga.PopulationSize, currentFitness)

		if currentFitness > bestPortfolioFitness {
			bestPortfolioFitness = currentFitness
			// Deep clone the weights to avoid mutation in further generations
			bestWeights = genetic_weight.CloneWeights(topPortfolio[0].Weight)
			ga.BestPortfolio = make([]string, len(topPortfolio))
			for i, ind := range topPortfolio {
				ga.BestPortfolio[i] = ind.Symbol
			}
		}

		// Genetic operations: selection, crossover, and mutation
		parentWeights1, parentWeights2 := ga.SelectWeights(population)
		for i := 0; i < len(population); i++ {
			childWeights := ga.CrossoverWeights(&parentWeights1, &parentWeights2)
			ga.MutateWeights(childWeights)
			population[i].Weight = childWeights
		}
	}

	fmt.Printf("Optimization Complete. Best Portfolio Fitness: %f\n", bestPortfolioFitness)
	fmt.Println("Top 10 Stocks in the Best Portfolio:")
	for _, symbol := range ga.BestPortfolio {
		fmt.Println(symbol)
	}

	// Save best weights to a file
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
