/**
 * Author: Mitch Allen
 * File: weighted-index_test.go
 */

package pick

import (
	"testing"
)

func TestWeightedIndex(t *testing.T) {

	limit := 1000
	threshold := 150

	weights := []float32{0.50, 0.30, 0.20}

	for i := 0; i < 100; i++ {
		m := make(map[int]int)
		for j := 0; j < limit; j++ {
			m[WeightedIndex(weights)]++
		}
		if len(m) < len(weights) {
			t.Errorf("WeightedIndex() value didn't return all items at least once: %v", m)
			return
		}
		for _, count := range m {
			if count < threshold {
				t.Errorf("OneInt() values not very balanced: %v", m)
				return
			}
		}
	}
}

func TestWeightedIndexOutOfBounds(t *testing.T) {

	limit := 1000

	weights := []float32{0.1, 0.1, 0.1}

	for i := 0; i < 100; i++ {
		m := make(map[int]int)
		for j := 0; j < limit; j++ {
			m[WeightedIndex(weights)]++
		}
		if m[-1] == 0 {
			t.Errorf("WeightedIndex() should have returned -1 at leasr once: %v", m)
			return
		}
	}
}
