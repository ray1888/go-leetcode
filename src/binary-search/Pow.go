package binary_search

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func myPow(x float64, n int) float64 {
	// int >1 or < 0
	// xiaoshu
	if x == 0 {
		return 0
	} else if n == 0 || x == float64(1) {
		return 1
	} else if x == float64(-1) {
		k := abs(n)
		result := float64(-1)
		if k%2 == 0 {
			result = float64(1)
		}
		return result
	}
	i := 0
	var result float64 = x
	k := abs(n)
	for i < k-1 {
		result *= x
		i += 1
	}
	if n < 0 {
		result = float64(1) / result
	}
	return result
}
