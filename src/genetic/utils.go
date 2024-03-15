package genetic

import (
	"encoding/json"
	"fmt"
	"os"

	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

func SaveBestWeights(bestWeights *genetic_weight.Weight) {
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

func (ga *GA) InitPopulation() []*Individual {
	population := make([]*Individual, ga.PopulationSize)

	for i := 0; i < ga.PopulationSize; i++ {
		individual := &Individual{}
		w := genetic_weight.InitializeRandomWeight()
		individual.Weight = &w
		individual.PortfolioScore = ga.EvaluateIndividual(individual, false)
		population[i] = individual
	}
	return population
}
