/**
 * Author: Mitch Allen
 * File: one-string.go
 */

package pick

// OneString returns a random item from a list
func OneString(list []string) string {
	return list[RandomIndex(len(list))]
}
