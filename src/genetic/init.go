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
	"github.com/spacecodewor/fmpcloud-go/objects"
)

type GA struct {
	MutationRate         float64
	PopulationSize       int
	Generations          int
	APIClient            *FMP.FMPAPI
	TickerFundamentals   map[string]*Individual
	BestPortfolio        []string
	mtx                  sync.Mutex
	EliteCount           int
	CrossoverRate        float64
	TournamentSize       int
	RiskFreeRate         float64
	AcceptableStagnation int
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
		if individual.Fundamentals == nil || individual.StockCandles == nil {
			fmt.Printf("Nil value for fundamentals or stock candles for ticker %s\n", ticker)
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

func NewGA(mutationRate float64, populationSize, generations, eliteCount int, crossoverRate float64, TournamentSize int, RiskFreeRate float64, AcceptableStagnation int, APIClient *FMP.FMPAPI) *GA {
	return &GA{
		MutationRate:         mutationRate,
		PopulationSize:       populationSize,
		Generations:          generations,
		APIClient:            APIClient,
		EliteCount:           eliteCount,
		CrossoverRate:        crossoverRate,
		TournamentSize:       TournamentSize,
		RiskFreeRate:         RiskFreeRate,
		AcceptableStagnation: AcceptableStagnation,
	}
}

func (ga *GA) NewRandomIndividual(symbol string) (*Individual, error) {
	Fundamentals, StockCandles, err := ga.APIClient.GetValuationInfo(symbol, "quarterly")
	if err != nil {
		return nil, fmt.Errorf("error getting company valuation info for %s: %v", symbol, err)
	}
	if Fundamentals == nil || StockCandles == nil {
		return nil, fmt.Errorf("error getting company valuation info for %s: nil value", symbol)
	}

	w := genetic_weight.InitializeRandomWeight()

	return &Individual{
		Symbol:       symbol,
		Fundamentals: Fundamentals,
		Weight:       &w,
		StockCandles: StockCandles,
	}, nil
}

func (ga *GA) TopPerformers(population []*Individual, tickerPopulation []string, count int) []string {
	// First, evaluate the weight scores for each ticker in the ticker population
	weightScores := make(map[string]float64)
	for _, ticker := range tickerPopulation {
		individual, exists := ga.TickerFundamentals[ticker]
		if !exists || individual == nil {
			fmt.Printf("Got nil after fetching fundamentals for %s\n", ticker)
			continue
		}
		weightScore := genetic_weight.CompositeWeightScore(individual.Fundamentals, population[0].Weight)
		weightScores[ticker] = weightScore
	}

	// Sort tickers by weight scores in descending order
	type kv struct {
		Key   string
		Value float64
	}
	var ss []kv
	for k, v := range weightScores {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	// Select top performers
	var topPerformers []string
	for _, kv := range ss[:count] {
		topPerformers = append(topPerformers, kv.Key)
	}

	return topPerformers
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
	var prevBestFitness float64
	var stagnantGenerations int

	for generation := 0; generation < ga.Generations; generation++ {
		fmt.Printf("Initializing Generation %d\n", generation)

		// Initialize population slice with the desired capacity.
		population := make([]*Individual, ga.PopulationSize)

		// Fill the population slice.
		for i := 0; i < ga.PopulationSize; i++ {
			w := genetic_weight.InitializeRandomWeight()
			individual := &Individual{
				Weight: &w,
			}
			population[i] = individual
		}

		// Evaluate fitness in parallel
		var wg sync.WaitGroup
		wg.Add(len(population))
		for _, individual := range population {
			go func(individual *Individual) {
				defer wg.Done()
				topPortfolio := ga.TopPerformers(population, TickerPopulation, 10)

				portfolioCandles := []*objects.StockDailyCandleList{}
				for _, ticker := range topPortfolio {
					stockInfo, exists := ga.TickerFundamentals[ticker]
					if !exists || stockInfo == nil {
						fmt.Printf("Got nil after fetching fundamentals for %s\n", ticker)
						continue
					}
					portfolioCandles = append(portfolioCandles, stockInfo.StockCandles)
				}

				individual.FundamentalScore = genetic_weight.CalculatePortfolioScore(portfolioCandles, ga.RiskFreeRate)
			}(individual)
		}
		wg.Wait()

		// Find the best portfolio in the current generation
		bestPortfolioInGeneration := ga.TopPerformers(population, TickerPopulation, 10)

		portfolioCandles := []*objects.StockDailyCandleList{}
		for _, ticker := range bestPortfolioInGeneration {
			stockInfo, exists := ga.TickerFundamentals[ticker]
			if !exists || stockInfo == nil {
				fmt.Printf("Got nil after fetching fundamentals for %s\n", ticker)
				continue
			}
			portfolioCandles = append(portfolioCandles, stockInfo.StockCandles)
		}

		fmt.Println("Calculating portfolio score")
		currentFitness := genetic_weight.CalculatePortfolioScore(portfolioCandles, ga.RiskFreeRate)
		fmt.Printf("Generation %d: Portfolio Fitness: %f\n", generation, currentFitness)
		mutex.Lock()
		if generation == 0 || currentFitness > bestPortfolioFitness {
			fmt.Printf("Best Portfolio Fitness Updated: %f\n", currentFitness)
			bestPortfolioFitness = currentFitness
			bestWeights = genetic_weight.CloneWeights(population[0].Weight)
			if bestWeights == nil {
				fmt.Println("Error cloning best weights")
				mutex.Unlock()
				continue
			}
			ga.BestPortfolio = bestPortfolioInGeneration
			fmt.Printf("Best Portfolio Updated: %v\n", ga.BestPortfolio)
		}
		if generation == 0 || currentFitness < worstPortfolioFitness {
			worstPortfolioFitness = currentFitness
			fmt.Printf("Worst Portfolio Updated: %f\n", worstPortfolioFitness)
		}
		mutex.Unlock()

		// Elitism: Keep the best individuals from the previous generation
		elitePopulation := make([]*Individual, ga.EliteCount)
		copy(elitePopulation, population[:ga.EliteCount])

		// Create a new population for the next generation
		newPopulation := make([]*Individual, 0, ga.PopulationSize)

		// Add the elite individuals to the new population
		newPopulation = append(newPopulation, elitePopulation...)
		fmt.Printf("Added %d elite individuals to the new population\n", len(elitePopulation))

		// Perform genetic operations (selection, crossover, mutation) to fill the remaining population
		for len(newPopulation) < ga.PopulationSize {
			parent1, parent2 := ga.SelectWeights(population)

			if rand.Float64() < ga.CrossoverRate {
				child := ga.CrossoverWeights(&parent1, &parent2)
				ga.MutateWeights(child)
				newPopulation = append(newPopulation, &Individual{Weight: child})
			} else {
				newPopulation = append(newPopulation, &Individual{Weight: &parent1}, &Individual{Weight: &parent2})
			}
		}
		fmt.Printf("New population size: %d\n", len(newPopulation))

		// Evaluate fitness of the new population in parallel
		var newWg sync.WaitGroup
		newWg.Add(len(newPopulation))
		for _, individual := range newPopulation {
			go func(individual *Individual) {
				defer newWg.Done()
				topPortfolio := ga.TopPerformers(newPopulation, TickerPopulation, 10)

				portfolioCandles := []*objects.StockDailyCandleList{}
				for _, ticker := range topPortfolio {
					stockInfo, exists := ga.TickerFundamentals[ticker]
					if !exists || stockInfo == nil {
						fmt.Printf("Got nil after fetching fundamentals for %s\n", ticker)
						continue
					}
					portfolioCandles = append(portfolioCandles, stockInfo.StockCandles)
				}

				individual.FundamentalScore = genetic_weight.CalculatePortfolioScore(portfolioCandles, ga.RiskFreeRate)
			}(individual)
		}
		newWg.Wait()

		// Adaptive mutation rate: Adjust the mutation rate based on the progress of the optimization
		if generation > 0 && generation%10 == 0 {
			improvementRatio := (bestPortfolioFitness - prevBestFitness) / prevBestFitness
			if improvementRatio < 0.01 {
				ga.MutationRate *= 1.1 // Increase mutation rate if improvement is slow
			} else if improvementRatio > 0.05 {
				ga.MutationRate *= 0.9 // Decrease mutation rate if improvement is fast
			}
			prevBestFitness = bestPortfolioFitness
		}

		// Catastrophic event: If the best fitness hasn't improved for a certain number of generations, introduce a catastrophic event
		if generation > 0 && bestPortfolioFitness == prevBestFitness {
			stagnantGenerations++
			if stagnantGenerations >= ga.AcceptableStagnation {
				fmt.Println("Evolution is stagnant. Introducing a catastrophic event.")
				catastrophicPopulation := make([]*Individual, 0, ga.PopulationSize/2)
				for i := 0; i < ga.PopulationSize/2; i++ {
					w := genetic_weight.InitializeRandomWeight()
					individual := &Individual{
						Weight: &w,
					}
					catastrophicPopulation = append(catastrophicPopulation, individual)
				}
				newPopulation = append(newPopulation, catastrophicPopulation...)
				ga.Generations += ga.Generations / 4
				stagnantGenerations = 0
			}
		} else {
			stagnantGenerations = 0
		}

		// Update mutation rate
		mutRateMin := (0.001 + 0.005) / 2
		mutRateMax := (0.01 + 0.02) / 2
		mutRate := mutRateMin + rand.Float64()*(mutRateMax-mutRateMin)
		mutRate = mutRate + ga.MutationRate // the last generation should have some effect on the mutation rate
		ga.MutationRate = mutRate / 2
		fmt.Printf("Adjusted mutation rate for next generation (%d): %f\n", generation+1, ga.MutationRate)
	}

	fmt.Printf("Optimization Complete. Best Portfolio Fitness: %f, Worst Portfolio Fitness: %f, Difference: %f\n", bestPortfolioFitness, worstPortfolioFitness, bestPortfolioFitness-worstPortfolioFitness)
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
