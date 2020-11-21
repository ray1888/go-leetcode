package stringProblem

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)

	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return result
	}

	pc := make([]int, 26)
	for _, item := range p {
		aa := int(item - 'a')
		pc[aa]++
	}

	left := 0
	right := 0
	for right < len(s) {
		if pc[s[right]-'a'] > 0 {
			pc[s[right]-'a']--
			right++
		} else {
			pc[s[left]-'a']++
			left++
		}
		if right-left == len(p) {
			result = append(result, left)
		}
	}
	return result
}
