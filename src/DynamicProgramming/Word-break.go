package DynamicProgramming

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return false
	}

	wordmap := make(map[string]bool)
	for _, word := range wordDict {
		wordmap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			_, exists := wordmap[s[j:i]]
			if exists && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
