package numbers

func isPerfectSquareBinarySearch(input int) bool {
	low := 1
	high := input
	for low < high {
		mid := (low + high) / 2
		square := mid * mid
		if square == input {
			return true
		} else if square > input {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func isPerfectSquareNewton(input int) bool {
	x := input
	for x*x > input {
		x = (x + input/x) / 2
	}
	return x*x == input
}
