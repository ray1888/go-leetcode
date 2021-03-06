package stringProblem

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	result := make([]int, 0)
	result = append(result, 0)
	tmp := 1
	for i := 0; i < n; i++ {
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, tmp+result[j])
		}
		tmp <<= 1
	}
	return result
}
