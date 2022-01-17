/**
 * Author: Mitch Allen
 * File: init.go
 */

package pick

import (
	"math/rand"
	"time"
)

// Seed rand
func init() {
	rand.Seed(time.Now().UnixNano())
}
