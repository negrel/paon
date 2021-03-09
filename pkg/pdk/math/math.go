package math

// Constrain the value between the given minimum and maximum.
func Constrain(value, min, max int) int {
	value = Max(value, min)
	value = Min(value, max)

	return value
}

// Min returns the smallest integer.
func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// Max returns the biggest integer.
func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
