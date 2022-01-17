/**
 * Author: Mitch Allen
 * File: weighted-index.go
 */

package pick

import (
	"math/rand"
)

// Returns a random weighted index
func WeightedIndex(weights []float32) int {
	rnd := rand.Float32()
	var lower float32 = 0.00
	for i, weight := range weights {
		upper := lower + weight
		if rnd >= lower && rnd < upper {
			return i
		}
		lower = upper
	}

	// Never reached 100% and random
	// number is out of bounds
	return -1
}
