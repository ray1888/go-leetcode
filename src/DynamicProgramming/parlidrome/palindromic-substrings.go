package parlidrome

func countSubstrings(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	length := len(s)
	counter := 0
	dp := make([][]bool, length)
	singleDis := make([]bool, length*length)

	for i := 0; i < length; i++ {
		// here dp mean s from index i to index j is Palindromic string with bool result
		dp[i], singleDis = singleDis[:length], singleDis[length:]
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
				counter++
			}
		}
	}

	return counter
}
