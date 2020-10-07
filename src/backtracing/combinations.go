package backtracing

func combineSub(n int, k int, start int, elem []int, result *[][]int) {
	if k == 0 {
		newElem := make([]int, len(elem))
		copy(newElem, elem)
		*result = append(*result, newElem)
	} else {
		for i := start; i <= n-k+1; i++ {
			elem = append(elem, i)
			combineSub(n, k-1, i+1, elem, result)
			elem = elem[:len(elem)-1]
		}
	}
}

func combine(n int, k int) [][]int {
	if k > n {
		return [][]int{}
	}
	result := make([][]int, 0)
	elem := make([]int, 0)
	combineSub(n, k, 1, elem, &result)
	return result
}
