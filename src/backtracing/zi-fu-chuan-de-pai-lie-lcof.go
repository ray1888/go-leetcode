package backtracing

func perMutation(s []byte, result *[]string, length int, start int) {
	if start == length {
		*result = append(*result, string(s))
		return
	}
	alphaMap := make(map[byte]bool)
	for i := start; i <= length; i++ {
		if _, ok := alphaMap[s[i]]; ok {
			continue
		}
		alphaMap[s[i]] = true
		s[start], s[i] = s[i], s[start]
		perMutation(s, result, length, start+1)
		s[start], s[i] = s[i], s[start]
	}
}

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	sa := make([]byte, 0)
	for _, item := range s {
		sa = append(sa, byte(item))
	}
	perMutation(sa, &result, len(s)-1, 0)
	return result
}
