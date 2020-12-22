package DynamicProgramming

import "sort"

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func longestStrChain(words []string) int {
	dp := make(map[string]int)

	w := ByLen(words)
	sort.Sort(w)
	maxValue := 1
	for _, word := range w {
		if _, ok := dp[word]; !ok {
			dp[word] = 1
		}
		for i := 0; i < len(word); i++ {
			newWord := word[:i] + word[i+1:]
			if val, ok := dp[newWord]; ok {
				dp[word] = max(val+1, dp[word])
			}
		}
		maxValue = max(dp[word], maxValue)
	}
	return maxValue
}
