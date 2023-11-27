package minmax

import "cmp"

// Constrain the value between the given minimum and maximum.
func Constrain[T cmp.Ordered](value, min, max T) T {
	value = Max(value, min)
	value = Min(value, max)

	return value
}

// Min returns the smallest integer.
func Min[T cmp.Ordered](nbrs ...T) T {
	min := nbrs[0]
	for _, nbr := range nbrs {
		if nbr < min {
			min = nbr
		}
	}

	return min
}

// Max returns the biggest integer.
func Max[T cmp.Ordered](nbrs ...T) T {
	max := nbrs[0]
	for _, nbr := range nbrs {
		if nbr > max {
			max = nbr
		}
	}

	return max
}
