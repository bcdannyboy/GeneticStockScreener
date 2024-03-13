package genetic

import (
	"math/rand"

	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

func (ga *GA) SelectWeights(population []*Individual) (genetic_weight.Weight, genetic_weight.Weight) {
	// Tournament selection
	tournamentSize := 5
	tournament1 := make([]*Individual, tournamentSize)
	tournament2 := make([]*Individual, tournamentSize)

	// Select individuals for the first tournament
	for i := 0; i < tournamentSize; i++ {
		randIndex := rand.Intn(len(population))
		tournament1[i] = population[randIndex]
	}

	// Select individuals for the second tournament
	for i := 0; i < tournamentSize; i++ {
		randIndex := rand.Intn(len(population))
		tournament2[i] = population[randIndex]
	}

	// Find the best individual in the first tournament
	bestIndividual1 := tournament1[0]
	for i := 1; i < tournamentSize; i++ {
		if tournament1[i].FundamentalScore > bestIndividual1.FundamentalScore {
			bestIndividual1 = tournament1[i]
		}
	}

	// Find the best individual in the second tournament
	bestIndividual2 := tournament2[0]
	for i := 1; i < tournamentSize; i++ {
		if tournament2[i].FundamentalScore > bestIndividual2.FundamentalScore {
			bestIndividual2 = tournament2[i]
		}
	}

	return *bestIndividual1.Weight, *bestIndividual2.Weight
}
