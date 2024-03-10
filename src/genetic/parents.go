package genetic

import (
	"math/rand"

	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

func (ga *GA) SelectWeights(population []*Individual) (genetic_weight.Weight, genetic_weight.Weight) {
	// Select first pair for the tournament
	randIndex1 := rand.Intn(len(population))
	randIndex2 := rand.Intn(len(population))
	for randIndex1 == randIndex2 {
		randIndex2 = rand.Intn(len(population)) // Ensure different indexes
	}

	// Determine winner of the first tournament
	weight1 := *population[randIndex1].Weight
	if population[randIndex1].FundamentalScore < population[randIndex2].FundamentalScore {
		weight1 = *population[randIndex2].Weight
	}

	// Select second pair for the tournament
	randIndex3 := rand.Intn(len(population))
	randIndex4 := rand.Intn(len(population))
	for randIndex3 == randIndex4 || randIndex3 == randIndex1 || randIndex3 == randIndex2 || randIndex4 == randIndex1 || randIndex4 == randIndex2 {
		randIndex3 = rand.Intn(len(population))
		randIndex4 = rand.Intn(len(population)) // Ensure different indexes and not repeating previous selections
	}

	// Determine winner of the second tournament
	weight2 := *population[randIndex3].Weight
	if population[randIndex3].FundamentalScore < population[randIndex4].FundamentalScore {
		weight2 = *population[randIndex4].Weight
	}

	return weight1, weight2
}
