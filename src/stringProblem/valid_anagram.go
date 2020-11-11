package stringProblem

func isAnagram(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	}
	if len(s) == 0 || len(t) == 0 || len(s) != len(t) {
		return false
	}
	nums := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		nums[s[i]-'a']++
		nums[t[i]-'a']--
	}
	for _, item := range nums {
		if item != 0 {
			return false
		}
	}
	return true
}
