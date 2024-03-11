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
	MutationFactor     float64
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
		if err != nil {
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

func NewGA(mutationRate, mutationFactor float64, populationSize, generations int, APIClient *FMP.FMPAPI) *GA {
	return &GA{
		MutationRate:   mutationRate,
		MutationFactor: mutationFactor,
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

	// Pre-fetch fundamentals for all tickers before the GA starts
	if err := ga.PreFetchFundamentals(TickerPopulation); err != nil {
		fmt.Println("Error pre-fetching fundamentals:", err)
		return
	}

	var bestPortfolioFitness float64
	var bestWeights *genetic_weight.Weight

	for generation := 0; generation < ga.Generations; generation++ {
		fmt.Printf("Initializing Generation %d\n", generation)

		// Generate the population for this generation using pre-fetched data
		population := make([]*Individual, ga.PopulationSize)
		for i := range population {
			tickerIndex := rand.Intn(len(TickerPopulation))
			ticker := TickerPopulation[tickerIndex]
			individual := ga.TickerFundamentals[ticker]

			// Clone the individual to work on a copy
			clonedIndividual := &Individual{
				Symbol:           individual.Symbol,
				PriceChangeScore: individual.PriceChangeScore,
				Fundamentals:     individual.Fundamentals,
				Weight:           genetic_weight.CloneWeights(individual.Weight), // Assume CloneWeights is properly implemented
			}

			population[i] = clonedIndividual
		}

		// Evaluate each individual's fitness based on their current weight
		for _, individual := range population {
			individual.FundamentalScore = genetic_weight.CompositeWeightScore(individual.Fundamentals, individual.Weight)
		}

		// Select the top performers to form a new population
		topPortfolio := topPerformers(population, 10)
		currentFitness := ga.EvaluatePortfolioFitness(topPortfolio)
		fmt.Printf("Generation %d: Portfolio Fitness: %f\n", generation, currentFitness)

		if currentFitness > bestPortfolioFitness {
			bestPortfolioFitness = currentFitness
			// Clone the best weight for future generations
			bestWeights = genetic_weight.CloneWeights(topPortfolio[0].Weight)

			// Save symbols of the top performers
			ga.BestPortfolio = make([]string, len(topPortfolio))
			for i, individual := range topPortfolio {
				ga.BestPortfolio[i] = individual.Symbol
			}
		}

		// Prepare for the next generation: selection, crossover, and mutation
		for i := 0; i < len(population); i++ {
			// Selection: Assume SelectWeights randomly selects two parents from the top performers
			parent1, parent2 := ga.SelectWeights(topPortfolio)
			// Crossover: Combine the selected parents to create a child
			child := ga.CrossoverWeights(&parent1, &parent2)
			// Mutation: Introduce random changes to the child's weights
			ga.MutateWeights(child)

			// Update the individual's weight with the newly generated child's weights
			population[i].Weight = child
		}
	}

	fmt.Printf("Optimization Complete. Best Portfolio Fitness: %f\n", bestPortfolioFitness)
	fmt.Println("Top 10 Stocks in the Best Portfolio:")
	for _, symbol := range ga.BestPortfolio {
		fmt.Println(symbol)
	}

	// Save the best weights to a file for future reference
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
