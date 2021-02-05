package parlidrome

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	length := len(s)
	dp := make([][]bool, length)

	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	start := 0
	maxLength := 0

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j] && dp[i+1][j-1])
			}
			if dp[i][j] && (j-i+1 > maxLength) {
				start = i
				maxLength = j - i + 1
			}
		}
	}
	return s[start : start+maxLength]
}
