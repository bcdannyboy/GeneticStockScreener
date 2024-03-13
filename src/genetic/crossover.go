package genetic

import (
	"math/rand"
	"reflect"

	genetic_weight "github.com/bcdannyboy/GeneticStockScreener/src/genetic/weight"
)

func (ga *GA) CrossoverWeights(parent1, parent2 *genetic_weight.Weight) *genetic_weight.Weight {
	child := &genetic_weight.Weight{}
	childValue := reflect.ValueOf(child).Elem()
	parent1Value := reflect.ValueOf(parent1).Elem()
	parent2Value := reflect.ValueOf(parent2).Elem()

	ga.crossoverStruct(childValue, parent1Value, parent2Value)

	return child
}

func (ga *GA) crossoverStruct(child, parent1, parent2 reflect.Value) {
	for i := 0; i < child.NumField(); i++ {
		childField := child.Field(i)
		parent1Field := parent1.Field(i)
		parent2Field := parent2.Field(i)

		switch childField.Kind() {
		case reflect.Float64:
			if rand.Float64() < 0.5 {
				childField.Set(parent1Field)
			} else {
				childField.Set(parent2Field)
			}
		case reflect.Struct:
			ga.crossoverStruct(childField, parent1Field, parent2Field)
		case reflect.Slice:
			ga.crossoverSlice(childField, parent1Field, parent2Field)
		}
	}
}

func (ga *GA) crossoverSlice(child, parent1, parent2 reflect.Value) {
	for i := 0; i < child.Len(); i++ {
		childElement := child.Index(i)
		parent1Element := parent1.Index(i)
		parent2Element := parent2.Index(i)

		if rand.Float64() < ga.CrossoverRate {
			childElement.Set(parent1Element)
		} else {
			childElement.Set(parent2Element)
		}
	}
}
