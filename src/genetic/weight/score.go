package genetic_weight

import (
	"log"
	"math"
	"reflect"

	"github.com/bcdannyboy/GeneticStockScreener/src/FMP"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func GetPriceChangeScore(PriceChange []objects.StockPriceChange) float64 {
	Sum := float64(0)
	Tot := 0

	for _, pc := range PriceChange {
		// double the weight of negative returns to penalize them
		if pc.OneD < 0 {
			pc.OneD *= 2
		} else if pc.FiveD < 0 {
			pc.FiveD *= 2
		} else if pc.OneM < 0 {
			pc.OneM *= 2
		} else if pc.ThreeM < 0 {
			pc.ThreeM *= 2
		} else if pc.SixM < 0 {
			pc.SixM *= 2
		} else if pc.Ytd < 0 {
			pc.Ytd *= 2
		} else if pc.OneY < 0 {
			pc.OneY *= 2
		} else if pc.FiveY < 0 {
			pc.FiveY *= 2
		} else if pc.TenY < 0 {
			pc.TenY *= 2
		}

		// penalize stocks with less data (newer enterprises)
		if pc.OneY == 0 {
			pc.OneY = -1
		}
		if pc.FiveY == 0 {
			pc.FiveY = -1
		}
		if pc.TenY == 0 {
			pc.TenY = -1
		}

		Sum += pc.OneD / (24)
		Sum += pc.FiveD / (24 * 5)
		Sum += pc.OneM / (24 * 5 * 4)
		Sum += pc.ThreeM / (24 * 5 * 4 * 3)
		Sum += pc.SixM / (24 * 5 * 4 * 6)
		Sum += pc.Ytd / (24 * 5 * 4 * 6) // YTD = 6 months
		Sum += pc.OneY / (24 * 5 * 4 * 12)
		Sum += pc.FiveY / (24 * 5 * 4 * 12 * 5)
		Sum += pc.TenY / (24 * 5 * 4 * 12 * 10)
		// Sum += pc.Max / (24 * (240 * 20))
		Tot += 9
	}

	// should be average return / hr across all timeframes for all stocks in the portfolio
	return Sum / float64(Tot)
}

// CompositeWeightScore calculates a weighted score for a company's valuation info based on given weights.
func CompositeWeightScore(companyInfo *FMP.CompanyValuationInfo, weights *Weight) float64 {
	if companyInfo == nil {
		log.Fatal("Company info is nil")
	}
	if weights == nil {
		log.Fatal("Weights are nil")
	}
	// Validate input types.
	if reflect.TypeOf(companyInfo).Kind() != reflect.Ptr || reflect.TypeOf(weights).Kind() != reflect.Ptr {
		log.Fatal("Input should be pointers to struct")
	}

	valInfo := reflect.Indirect(reflect.ValueOf(companyInfo))
	valWeights := reflect.Indirect(reflect.ValueOf(weights))

	// Ensure we're dealing with structs.
	if valInfo.Kind() != reflect.Struct || valWeights.Kind() != reflect.Struct {
		log.Fatal("Expecting struct types")
	}

	totalScore, totalWeight := processFields(valInfo, valWeights)

	// Avoid division by zero.
	if totalWeight != 0 {
		return totalScore / totalWeight // Normalize by totalWeight.
	}

	return 0
}

// processFields iterates over each field in the valuation info struct and calculates a weighted score.
func processFields(valInfo, valWeights reflect.Value) (totalScore float64, totalWeight float64) {
	for i := 0; i < valInfo.NumField(); i++ {
		fieldInfo := valInfo.Field(i)
		fieldType := valInfo.Type().Field(i)
		weightField, found := valWeights.Type().FieldByName(fieldType.Name)

		// If a corresponding weight field is found, calculate the score for the field.
		if found {
			fieldWeights := valWeights.FieldByName(weightField.Name)
			itemScore, itemWeight := calculateFieldScore(fieldInfo, fieldWeights)
			totalScore += itemScore
			totalWeight += itemWeight
		} else {
			// Log a warning if no corresponding weight field is found.
			log.Printf("Warning: Missing weight for '%v'. Setting to default.", fieldType.Name)
		}
	}

	return totalScore, totalWeight
}

// calculateFieldScore computes the score for a given field, adjusting for the field's weight.
func calculateFieldScore(dataField, weightField reflect.Value) (score float64, weight float64) {
	// Only proceed if both fields are struct types; otherwise, return zero values.
	if dataField.Kind() != reflect.Struct || weightField.Kind() != reflect.Struct {
		return 0, 0
	}

	// Iterate over fields in the data struct.
	for i := 0; i < dataField.NumField(); i++ {
		field := dataField.Field(i)

		// Ensure the field is of type float64 and a corresponding weight exists.
		if field.Kind() == reflect.Float64 {
			// Use the same index for weights, assuming alignment.
			if i < weightField.NumField() {
				weightValue := weightField.Field(i).Float()

				// Calculate the weighted score for the field.
				dataValue := field.Float()
				score += dataValue * weightValue
				weight += math.Abs(weightValue)
			}
		}
	}

	return score, weight
}
