package DynamicProgramming

func LongestCountSubstrings(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	length := len(s)
	maxium := 0
	dp := make([][]bool, length)
	single_dis := make([]bool, length*length)

	for i := 0; i < length; i++ {
		// here dp mean s from index i to index j is Palindromic string with bool result
		dp[i], single_dis = single_dis[:length], single_dis[length:]
	}

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j] && dp[i+1][j-1])
			}
			if dp[i][j] {
				maxium = max(maxium, j-i+1)
			}
		}
	}

	return maxium
}

func Longestexpand(s string, left, right int) int {
	length := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		length = (right - left + 1)
		left -= 1
		right += 1
	}
	return length
}

func LongestSubstringsExpand(s string) int {
	if len(s) == 0 || s == "" {
		return 0
	}
	maxium := 0
	i_longest := 0
	for i := 0; i < len(s); i++ {
		i_longest = max(Longestexpand(s, i, i), Longestexpand(s, i, i+1))
		maxium = max(maxium, i_longest)
	}
	return maxium
}
