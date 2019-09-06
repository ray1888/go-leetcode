package binary_search

func mySqrt(x int) int {
	start := 1
	end := x
	for start <= end {
		mid := start + (end-start)/2
		square_mid := mid * mid
		if square_mid > x {
			end = mid - 1
		} else if square_mid < x {
			start = mid + 1
		} else {
			return mid
		}
	}
	return end
}
