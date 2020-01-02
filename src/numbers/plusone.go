package numbers

func PlusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	p := len(digits) - 1
	for p >= 0 && digits[p] == 9 {
		digits[p] = 0
		p -= 1
	}
	if p == -1 {
		result := make([]int, len(digits)+1)
		result[0] = 1
		return result
	} else {
		digits[p] += 1
		return digits
	}
}
