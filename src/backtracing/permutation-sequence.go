package backtracing

import "strconv"

var (
	factorial []int
	K         int
	used      []bool
)

func calfatorial(n int) {
	factorial = make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}
}

/*

	为了满足这个剪枝的需求，

	实际上这里添加了一个判断，
	就是把第一个元素去掉之后，能够组合的数和第k个数做对比，如果小于，则直接跳过，进入下一个数为第一个数的循环
*/

func getPermutation2(n int, k int) string {
	if n == 0 {
		return ""
	}
	K = k
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	calfatorial(n)
	used = make([]bool, n+1)
	var b []int
	getPermutationRec2(0, n, &b)
	s := ""
	for _, item := range b {
		a := strconv.Itoa(item)
		s += a
	}
	return s
}

func getPermutationRec2(start int, n int, path *[]int) {
	if start == n {
		return
	}
	cnt := factorial[n-1-start]
	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}
		if cnt < K {
			K -= cnt
			continue
		}
		*path = append(*path, i)
		used[i] = true
		getPermutationRec2(start+1, n, path)
		return
	}
}
