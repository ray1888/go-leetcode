package numbers

func lexicalOrder(n int) []int {
	ans := []int{}
	var dfs func(m int)
	dfs = func(m int) {
		if m > n {
			return
		}
		for i := 0; i < 10; i++ {
			if m+i <= n && m+i != 0 {
				ans = append(ans, m+i)
				dfs(10 * (m + i))
			}
		}
	}
	dfs(0)
	return ans
}
