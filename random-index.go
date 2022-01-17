/**
 * Author: Mitch Allen
 * File: random-index.go
 */

package pick

import (
	"math"
	"math/rand"
)

// RandomIndex returns a random index based on the length of list
func RandomIndex(len int) int {
	return int(math.Floor(rand.Float64() * float64(len)))
}
