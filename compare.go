// Package arithmeticutil contains utility functions for working with arithmetic operation with numbers
package arithmeticutil

// Min returns min between two arguments
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
