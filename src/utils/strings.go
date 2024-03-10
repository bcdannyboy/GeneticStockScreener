package utils

import (
	"math/rand"
	"time"
)

func GetRandomFromStrArray(A []string, n int) []string {
	// Seed the random number generator to ensure different outcomes on each run
	rand.Seed(time.Now().UnixNano())

	// Create a slice to store the randomly selected elements
	var randomSelection []string

	// Ensure n does not exceed the length of A to avoid infinite loops
	if n > len(A) {
		n = len(A)
	}

	// Use a map to track which indices have already been selected
	selectedIndices := make(map[int]bool)

	// Keep selecting random elements until we have n unique selections
	for len(randomSelection) < n {
		index := rand.Intn(len(A))
		// Check if the index was already selected
		if _, exists := selectedIndices[index]; !exists {
			randomSelection = append(randomSelection, A[index])
			selectedIndices[index] = true
		}
	}

	return randomSelection
}
