/**
 * Author: Mitch Allen
 * File: random-index.go
 */

package pick

import (
	"math"
	"math/rand"
	"time"
)

// RandomIndex returns a random index based on the length of list
func RandomIndex(len int) int {
	rand.Seed(time.Now().UnixNano())
	return int(math.Floor(rand.Float64() * float64(len)))
}
