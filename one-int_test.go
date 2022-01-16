/**
 * Author: Mitch Allen
 * File: one-int_test.go
 */

package pick

import (
	"testing"
)

func TestOneInt(t *testing.T) {

	limit := 1000
	threshold := 250

	list := []int{20, 10, 30}

	for i := 0; i < 100; i++ {
		m := make(map[int]int)
		for j := 0; j < limit; j++ {
			m[OneInt(list)]++
		}
		for _, val := range list {
			if m[val] == 0 {
				t.Errorf("OneInt() value didn't return all items at least once: %v", m)
				return
			}
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
