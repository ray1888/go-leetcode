package DynamicProgramming

func stringPartition(dp [][]bool, s string, start, end int, res *[][]string, tmp []string) {
	if start == end {
		t := make([]string, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	for j := start; j < end; j++ {
		if dp[start][j] {
			tmp = append(tmp, s[start:j+1])
			// add one to do recursive call is important
			stringPartition(dp, s, j+1, end, res, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func partition(s string) [][]string {
	if s == "" {
		return [][]string{}
	}
	// dp[i][j] mean string from i to j isparlindrome
	length := len(s)
	dp := make([][]bool, length*length)
	single_dis_dp := make([]bool, length*length)
	for i := 0; i < length; i++ {
		dp[i], single_dis_dp = single_dis_dp[:length], single_dis_dp[length:]
	}

	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (i-j < 2 || dp[i-1][j+1]) {
				dp[j][i] = true
			}
		}
	}
	res := make([][]string, 0, 50)
	tmp := make([]string, 0, 10)
	stringPartition(dp, s, 0, length, &res, tmp)
	return res
}
