package math

import "github.com/negrel/debuggo/pkg/assert"

// Constrain the value between the given minimum and maximum.
func Constrain(value, min, max int) int {
	value = Max(value, min)
	value = Min(value, max)

	return value
}

// Min returns the smallest integer.
func Min(nbrs ...int) int {
	assert.Greater(0, len(nbrs))

	min := nbrs[0]
	for _, nbr := range nbrs {
		if nbr < min {
			min = nbr
		}
	}

	return min
}

// Max returns the biggest integer.
func Max(nbrs ...int) int {
	assert.Greater(0, len(nbrs))

	max := nbrs[0]
	for _, nbr := range nbrs {
		if nbr > max {
			max = nbr
		}
	}

	return max
}
