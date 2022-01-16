/**
 * Author: Mitch Allen
 * File: oneint.go
 */

package pick

import (
	"math"
	"math/rand"
	"time"
)

// OneInt returns a random item from a list
func OneInt(list []int) int {
	rand.Seed(time.Now().UnixNano())
	i := int(math.Floor(rand.Float64() * float64(len(list))))
	return list[i]
}
