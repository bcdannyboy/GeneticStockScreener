package genetic_weight

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"sync"

	"github.com/bcdannyboy/GeneticStockScreener/src/FMP"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

var cache = make(map[string]float64)
var mu sync.Mutex

func CalculatePortfolioScore(portfolio []*objects.StockDailyCandleList, riskFreeRate float64) float64 {
	dailyRiskFreeRate := math.Pow((riskFreeRate+1), 1.0/365) - 1

	if len(cache) > 100000 {
		// Clear the cache if it gets too large
		cache = make(map[string]float64)
	}

	// Generate a cache key
	cacheKey := generateCacheKey(portfolio, riskFreeRate)

	// Check if the result is cached
	mu.Lock()
	if score, found := cache[cacheKey]; found {
		mu.Unlock()
		return score
	}
	mu.Unlock()

	var sharpeRatios []float64

	for _, stock := range portfolio {
		var returns []float64

		for i := 1; i < len(stock.Historical); i++ {
			currentPrice := stock.Historical[i].Close
			previousPrice := stock.Historical[i-1].Close
			returnValue := (currentPrice - previousPrice) / previousPrice
			returns = append(returns, returnValue)
		}

		averageReturn := calculateAverage(returns)
		stdDev := calculateStandardDeviation(returns)

		if stdDev != 0 {
			sharpeRatio := (averageReturn - dailyRiskFreeRate) / stdDev
			sharpeRatios = append(sharpeRatios, sharpeRatio)
		}
	}

	averageSharpeRatio := calculateAverage(sharpeRatios)

	// Cache the result before returning
	mu.Lock()
	cache[cacheKey] = averageSharpeRatio
	mu.Unlock()

	return averageSharpeRatio
}

func generateCacheKey(portfolio []*objects.StockDailyCandleList, riskFreeRate float64) string {
	// Use the symbols and the risk-free rate to generate a unique key
	symbols := ""
	for _, stock := range portfolio {
		symbols += stock.Symbol + ","
	}
	return fmt.Sprintf("%s%f", symbols, riskFreeRate)
}

func calculateAverage(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	sum := 0.0
	for _, value := range values {
		sum += value
	}
	return sum / float64(len(values))
}

func calculateStandardDeviation(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	average := calculateAverage(values)
	sum := 0.0
	for _, value := range values {
		sum += math.Pow(value-average, 2)
	}
	variance := sum / float64(len(values))
	return math.Sqrt(variance)
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
				// Multiply the field value by its corresponding weight.
				score += field.Float() * weightValue

				// Add the absolute value of the weight to the total weight.
				weight += math.Abs(weightValue)
			}
		}
	}

	return score, weight
}
