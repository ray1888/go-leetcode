package numbers

func singleNumber(nums []int) int {
	var one, two, three int
	for _, num := range nums {
		two |= one & num
		one ^= num
		three = one & two
		one -= three // ones &= ~threes;
		two -= three // twos &= ~threes;
	}
	return one
}
