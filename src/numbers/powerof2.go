package numbers

func PowerOf2(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n = n / 2
	}
	return n == 1
}

func PowerOf2Math(n int) bool {
	return (n > 0) && (n&(n-1)) == 0
}
