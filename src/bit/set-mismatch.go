package bit

func findErrorNums(nums []int) []int {
	var xor, xor0, xor1 int

	for index, item := range nums {
		xor = xor ^ item ^ (index + 1)
	}

	rightMostBit := xor & (-xor)
	for _, item := range nums {
		if (item & rightMostBit) != 0 {
			xor1 ^= item
		} else {
			xor0 ^= item
		}
	}
	for index := range nums {
		i := index + 1
		if (i & rightMostBit) != 0 {
			xor1 ^= i
		} else {
			xor0 ^= i
		}
	}

	for _, item := range nums {
		if item == xor0 {
			return []int{xor0, xor1}
		}
	}
	return []int{xor1, xor0}
}
