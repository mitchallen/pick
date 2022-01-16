/**
 * Author: Mitch Allen
 * File: one-int.go
 */

package pick

// OneInt returns a random item from a list
func OneInt(list []int) int {
	return list[RandomIndex(len(list))]
}
