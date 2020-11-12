package numbers

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func reachNumber(target int) int {
	t := abs(target)
	n := 0
	sum := 0
	for sum < t || (((sum - t) & 1) == 1) {
		n++
		sum += n
	}
	return n
}
