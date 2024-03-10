package genetic

import (
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
	if rand.Float64() < 0.05 { // Assuming a 5% mutation rate
		mutationFactor := 1 + (rand.Float64()*0.2 - 0.1) // ±10%
		mutatedValue := field.Float() * mutationFactor
		field.SetFloat(mutatedValue)
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
