package stringProblem

func convertToTitle(n int) string {
	s := make([]byte, 0)
	for n > 0 {
		n--
		var c byte
		c = byte(n%26) + 'A'
		s = append(s, c)
		n /= 26
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
