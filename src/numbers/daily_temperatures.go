package numbers

func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return T
	}
	stack := make([]int, 0, len(T))
	result := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) != 0 && T[i] > T[stack[len(stack)-1]] {
			// fmt.Println(stack)
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return result
}
