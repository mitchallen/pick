/**
 * Author: Mitch Allen
 * File: oneint_test.go
 */

package pick

import (
	"sort"
	"testing"
)

func TestOneInt(t *testing.T) {
	list := []int{10, 20, 30}
	for i := 0; i < 100; i++ {
		if got := OneInt(list); sort.SearchInts(list, got) >= len(list) {
			t.Errorf("OneInt() = %d, didn't return value from list", got)
		}
	}
}

func TestOneIntAverage(t *testing.T) {

	limit := 1000
	threshold := 250

	list := []int{10, 20, 30}

	for i := 0; i < 100; i++ {
		m := make(map[int]int)
		for j := 0; j < limit; j++ {
			m[OneInt(list)]++
		}
		if len(m) < len(list) {
			t.Errorf("OneInt() value didn't return all items at least once: %v", m)
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
