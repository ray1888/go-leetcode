package roman

func romanToInt(s string) int {
	basic := make(map[uint8]int)
	basic['I'] = 1
	basic['V'] = 5
	basic['X'] = 10
	basic['L'] = 50
	basic['C'] = 100
	basic['D'] = 500
	basic['M'] = 1000

	result := 0
	result += basic[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if basic[s[i]] < basic[s[i+1]] {
			result -= basic[s[i]]
		} else {
			result += basic[s[i]]
		}
	}
	return result
}
