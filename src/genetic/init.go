package genetic

import (
	"fmt"
	"math"
	"math/rand"
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
	TickerFundamentals   map[string]*FMP.CompanyValuationInfo
	TickerCandles        map[string]*objects.StockDailyCandleList
	EliteCount           int
	CrossoverRate        float64
	TournamentSize       int
	RiskFreeRate         float64
	AcceptableStagnation int
	NumWorkers           int
}

func NewGA(mutationRate float64, populationSize, generations, eliteCount int, crossoverRate float64, TournamentSize int, RiskFreeRate float64, AcceptableStagnation int, APIClient *FMP.FMPAPI, TickerFundamentals map[string]*FMP.CompanyValuationInfo, TickerCandles map[string]*objects.StockDailyCandleList, numWorkers int) *GA {
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
		TickerFundamentals:   TickerFundamentals,
		TickerCandles:        TickerCandles,
		NumWorkers:           numWorkers,
	}
}

func (ga *GA) RunGeneticAlgorithm() (*genetic_weight.Weight, float64, float64, float64) {
	population := ga.InitPopulation()
	bestPortfolioScore := math.Inf(-1)
	lastBest := math.Inf(-1)
	worstPortfolioScore := math.Inf(1)
	var topWeight *genetic_weight.Weight
	stagnation := 0
	secondaryStagnation := 0

	for i := 0; i < ga.Generations; i++ {
		rand.Seed(time.Now().UnixNano()) // Seed for randomness in each generation
		pop, best, worst, topW := ga.EvolvePopulation(population)

		population = pop

		if best >= lastBest {
			if best > bestPortfolioScore {
				fmt.Println("New best portfolio score found")
				bestPortfolioScore = best
				topWeight = topW
				secondaryStagnation = 0
			} else {
				secondaryStagnation++
			}
			stagnation = 0 // Reset stagnation
		} else {
			fmt.Printf("Stagnation detected (Best: %f, Last Best: %f)\n", best, lastBest)
			stagnation++ // Increment stagnation
		}

		if worst < worstPortfolioScore {
			worstPortfolioScore = worst
		}

		// Introduce a random factor in mutation and crossover adjustments
		ga.MutationRate = 0.01 + rand.Float64()*0.1 + (rand.Float64() * float64(stagnation) / float64(ga.AcceptableStagnation))
		ga.CrossoverRate = 0.8 + rand.Float64()*0.1 + (rand.Float64() * float64(stagnation) / float64(ga.AcceptableStagnation))

		fmt.Printf("Generation (%d/%d): Generation's best: %f, last Best: %f, Generation's Worst: %f, Total Best: %f, Total Worst: %f, Mutation Rate: %f, Crossover Rate: %f, Stagnation: %d/%d, TopRateStagnation: %d/%d\n",
			i, ga.Generations, best, lastBest, worst, bestPortfolioScore, worstPortfolioScore, ga.MutationRate, ga.CrossoverRate, stagnation, ga.AcceptableStagnation, secondaryStagnation, int(float64(ga.Generations)*float64(0.75)))

		if stagnation >= ga.AcceptableStagnation || float64(secondaryStagnation) >= (float64(ga.Generations)*float64(0.75)) { // if stagnation of best solution happens for over 60% of population size or stagnation of the local solutions happens for ga.AcceptableStagnation generations
			fmt.Println("Stagnation detected, introducing more diversity")
			ga.introduceDiversity(population)
			stagnation = 0
			ga.Generations += ga.Generations / 4
			ga.AcceptableStagnation += ga.AcceptableStagnation / 4
			secondaryStagnation = secondaryStagnation / 2
		}

		lastBest = best
	}

	return topWeight, bestPortfolioScore, worstPortfolioScore, bestPortfolioScore / worstPortfolioScore
}

func (ga *GA) introduceDiversity(population []*Individual) {
	var wg sync.WaitGroup
	wg.Add(len(population))

	for i := range population {
		go func(i int) {
			defer wg.Done()

			minRep := 0.1
			maxRep := 0.25
			repChance := minRep + rand.Float64()*(maxRep-minRep)

			if rand.Float64() < 0.5 { // Apply mutation to 50% of the population
				ga.MutateWeights(population[i].Weight)
			} else if rand.Float64() < repChance { // Completely replace between 10% and 33% of the population
				randW := genetic_weight.InitializeRandomWeight()
				population[i] = &Individual{Weight: &randW, PortfolioScore: ga.EvaluateIndividual(&Individual{Weight: &randW}, false)}
			}
		}(i)
	}

	wg.Wait()
}

func (ga *GA) generateIndividual(nonElites []*Individual) *Individual {
	minRandIndividualChance := 0.05
	maxRandIndividualChance := 0.25
	randIndividualChance := minRandIndividualChance + rand.Float64()*(maxRandIndividualChance-minRandIndividualChance)

	if rand.Float64() > randIndividualChance {
		// Selection
		parent1, parent2 := ga.SelectWeights(nonElites)

		// Crossover
		childWeights := ga.CrossoverWeights(&parent1, &parent2)

		// Mutation
		ga.MutateWeights(childWeights)

		// Creating a new individual
		child := &Individual{Weight: childWeights}
		child.PortfolioScore = ga.EvaluateIndividual(child, false)

		return child
	} else {
		// Randomly generate a new individual
		randW := genetic_weight.InitializeRandomWeight()
		individual := &Individual{
			Weight: &randW,
		}
		individual.PortfolioScore = ga.EvaluateIndividual(individual, false)

		return individual
	}
}

func (ga *GA) EvolvePopulation(population []*Individual) ([]*Individual, float64, float64, *genetic_weight.Weight) {
	newPopulation := make([]*Individual, 0, ga.PopulationSize)
	bestPortfolioScore := math.Inf(-1)
	worstPortfolioScore := math.Inf(1)
	var topWeight *genetic_weight.Weight

	// Elite selection: directly pass the best individuals to the next generation
	elites, nonElites := ga.selectElites(population, ga.EliteCount)
	newPopulation = append(newPopulation, elites...)

	results := make(chan *Individual, ga.PopulationSize-len(elites))
	var wg sync.WaitGroup

	numToGenerate := ga.PopulationSize - len(elites)
	numWorkers := ga.NumWorkers
	if numWorkers > numToGenerate {
		numWorkers = numToGenerate
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(nonElites []*Individual) {
			defer wg.Done()
			for j := 0; j < numToGenerate/numWorkers; j++ {
				child := ga.generateIndividual(nonElites)
				results <- child
			}
		}(nonElites)
	}

	wg.Wait()
	close(results)

	for child := range results {
		newPopulation = append(newPopulation, child)
		if child.PortfolioScore > bestPortfolioScore {
			bestPortfolioScore = child.PortfolioScore
			topWeight = child.Weight
		}
		if child.PortfolioScore < worstPortfolioScore {
			worstPortfolioScore = child.PortfolioScore
		}
	}

	return newPopulation, bestPortfolioScore, worstPortfolioScore, topWeight
}

func (ga *GA) EvaluateIndividual(individual *Individual, printTickers bool) float64 {
	rand.Seed(time.Now().UnixNano())

	// Calculate fundamentals score for each ticker
	tickerScores := make([]TickerScore, len(ga.TickerFundamentals))
	var wg sync.WaitGroup
	wg.Add(len(ga.TickerFundamentals))

	i := 0
	for ticker, valuationInfo := range ga.TickerFundamentals {
		go func(ticker string, valuationInfo *FMP.CompanyValuationInfo, i int) {
			defer wg.Done()
			score := genetic_weight.CompositeWeightScore(valuationInfo, individual.Weight)
			tickerScores[i] = TickerScore{Ticker: ticker, Score: score}
		}(ticker, valuationInfo, i)
		i++
	}

	wg.Wait()

	// Sort tickers based on their score in descending order
	sort.Slice(tickerScores, func(i, j int) bool {
		return tickerScores[i].Score > tickerScores[j].Score
	})

	// Determine the number of tickers to select: rand(10, min(50, length of tickerScores))
	numTickers := 10
	if len(tickerScores) > 10 {
		numTickers += rand.Intn(min(50, len(tickerScores)) - 10)
	}
	if printTickers {
		numTickers = min(10, len(tickerScores))
	}

	selectedTickers := tickerScores[:numTickers]

	if printTickers {
		fmt.Println("Top 10 Tickers:")
		for _, ts := range selectedTickers {
			fmt.Printf("%s: %f\n", ts.Ticker, ts.Score)
		}
	}

	// Collect candle data for the selected tickers based on their fundamentals score
	var portfolio []*objects.StockDailyCandleList
	for _, ts := range selectedTickers {
		if candleData, exists := ga.TickerCandles[ts.Ticker]; exists {
			portfolio = append(portfolio, candleData)
		}
	}

	tickers := make([]string, len(portfolio))
	for i, ts := range selectedTickers {
		tickers[i] = ts.Ticker
	}

	// Calculate the portfolio score using financial metrics (Calmar and Sharpe ratios)
	portfolioScore := genetic_weight.CalculatePortfolioScore(portfolio, ga.RiskFreeRate)

	return portfolioScore
}

func (ga *GA) selectElites(population []*Individual, eliteCount int) ([]*Individual, []*Individual) {
	// Ensure that there are no nil individuals in the population
	removedNils := 0
	nonNilPopulation := make([]*Individual, 0, len(population))
	for _, individual := range population {
		if individual != nil {
			nonNilPopulation = append(nonNilPopulation, individual)
		} else {
			removedNils++
		}
	}

	if removedNils > 0 {
		fmt.Printf("Removed %d nil individuals from the population\n", removedNils)
	}
	// Sort the non-nil population
	sort.Slice(nonNilPopulation, func(i, j int) bool {
		return nonNilPopulation[i].PortfolioScore > nonNilPopulation[j].PortfolioScore
	})

	// Split into elites and non-elites, ensuring we don't exceed bounds
	if eliteCount > len(nonNilPopulation) {
		eliteCount = len(nonNilPopulation)
	}
	return nonNilPopulation[:eliteCount], nonNilPopulation[eliteCount:]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
