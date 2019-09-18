package DynamicProgramming

func isValidTwoDigits(a, b byte) bool {
	if (a == '1' && b <= '9') || (a == '2' && b <= '6') {
		return true
	} else {
		return false
	}
}

func numDecodingsrecur(s string, index int) int {
	if index == len(s) {
		return 1
	}
	if index > len(s) {
		return 0
	}
	result := 0
	if s[index] != '0' {
		result += numDecodingsrecur(s, index+1)
	}
	if (index+1 < len(s)) && isValidTwoDigits(s[index], s[index+1]) {
		result += numDecodingsrecur(s, index+2)
	}
	return result
}

func numDecodingsRecur(s string) int {
	return numDecodingsrecur(s, 0)
}

func numDecodingsDP(s string) int {
	if s == "" {
		return 0
	}

	dp := make([]int, len(s)+1)

	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i < (len(s) + 1); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if isValidTwoDigits(s[i-2], s[i-1]) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
