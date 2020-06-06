package helpers

// Min return the smallest integer.
func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// Max return the biggest integer.
func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
