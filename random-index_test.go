/**
 * Author: Mitch Allen
 * File: random-index_test.go
 */

package pick

import (
	"testing"
)

func TestRandomIndex(t *testing.T) {

	limit := 1000
	threshold := 250

	aLen := 3

	for i := 0; i < 100; i++ {
		m := make(map[int]int)
		for j := 0; j < limit; j++ {
			m[RandomIndex(aLen)]++
		}
		if len(m) < aLen {
			t.Errorf("RandomIndex() value didn't return all items at least once: %v", m)
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
