package genetic

import (
	"math/rand"
	"reflect"

	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

func (ga *GA) CrossoverWeights(weight1, weight2 *genetic_weight.Weight) *genetic_weight.Weight {
	offspringWeight := &genetic_weight.Weight{}

	val1 := reflect.ValueOf(weight1).Elem()
	val2 := reflect.ValueOf(weight2).Elem()
	valOffspring := reflect.ValueOf(offspringWeight).Elem()

	ga.performUniformCrossover(val1, val2, valOffspring)

	return offspringWeight
}

func (ga *GA) performUniformCrossover(val1, val2, valOffspring reflect.Value) {
	for i := 0; i < val1.NumField(); i++ {
		field1 := val1.Field(i)
		field2 := val2.Field(i)
		fieldOffspring := valOffspring.Field(i)

		switch field1.Kind() {
		case reflect.Float64:
			// For float64 fields, randomly choose the value from parent1 or parent2
			if rand.Float64() < 0.5 {
				fieldOffspring.SetFloat(field1.Float())
			} else {
				fieldOffspring.SetFloat(field2.Float())
			}
		case reflect.Struct:
			// For structs, recursively perform uniform crossover
			ga.performUniformCrossover(field1, field2, fieldOffspring)
		case reflect.Slice:
			// Handle slices, assuming they are slices of float64 for simplicity
			ga.handleSliceCrossover(field1, field2, fieldOffspring)
		}
	}
}

func (ga *GA) handleSliceCrossover(slice1, slice2, sliceOffspring reflect.Value) {
	// Determine the smaller length to avoid index out of range
	minLen := min(slice1.Len(), slice2.Len())
	newSlice := reflect.MakeSlice(slice1.Type(), minLen, minLen)

	for i := 0; i < minLen; i++ {
		if rand.Float64() < 0.5 {
			newSlice.Index(i).Set(slice1.Index(i))
		} else {
			newSlice.Index(i).Set(slice2.Index(i))
		}
	}

	sliceOffspring.Set(newSlice)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
