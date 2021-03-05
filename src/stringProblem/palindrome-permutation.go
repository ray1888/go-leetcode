package stringProblem

func canPermutePalindrome(s string) bool {
	alphaMap := make(map[rune]int)

	for _, item := range s {
		if _, ok := alphaMap[item]; !ok {
			alphaMap[item] = 0
		}
		alphaMap[item]++
	}

	cnt := 0
	for _, value := range alphaMap {
		if value%2 == 1 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}

	return true
}
