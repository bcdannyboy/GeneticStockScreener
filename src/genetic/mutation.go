package genetic

import (
	"math"
	"math/rand"
	"reflect"

	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

func (ga *GA) MutateWeights(weights *genetic_weight.Weight) {
	val := reflect.ValueOf(weights).Elem()
	ga.mutateStruct(val)
}

func (ga *GA) mutateStruct(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Float64:
			ga.mutateFloat64Field(field)
		case reflect.Struct:
			ga.mutateStruct(field)
		case reflect.Slice:
			ga.mutateSlice(field)
		}
	}
}

func (ga *GA) mutateFloat64Field(field reflect.Value) {
	if rand.Float64() < ga.MutationRate { // Dynamically use GA's mutation rate
		mutMin := 0.001
		mutMax := 0.5

		// Compute the mutation factor
		mutationFactor := mutMin + rand.Float64()*(mutMax-mutMin)

		// Randomly decide to increase or decrease
		changeDirection := rand.Intn(2)*2 - 1 // Results in -1 (decrease) or 1 (increase)

		// Calculate the mutation effect
		mutationEffect := float64(changeDirection) * mutationFactor

		// Apply mutation based on the current value to ensure it's within the [-1, 1] range
		currentValue := field.Float()
		mutatedValue := currentValue + mutationEffect*(1-math.Abs(currentValue))

		// Ensure mutatedValue does not exceed the bounds [-1, 1]
		if mutatedValue > 1 {
			mutatedValue = 1
		} else if mutatedValue < -1 {
			mutatedValue = -1
		}

		if field.CanSet() {
			field.SetFloat(mutatedValue)
		}
	}
}

func (ga *GA) mutateSlice(slice reflect.Value) {
	// Example mutation for slices of float64
	for i := 0; i < slice.Len(); i++ {
		element := slice.Index(i)
		if element.Kind() == reflect.Float64 {
			ga.mutateFloat64Field(element)
		}
	}
}
