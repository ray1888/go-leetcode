package numbers

func FactorilTrailingZeros(input int) int {
	if input == 0 {
		return 0
	}
	count := 0
	for input > 0 {
		input /= 5
		count += input
	}
	return count
}
