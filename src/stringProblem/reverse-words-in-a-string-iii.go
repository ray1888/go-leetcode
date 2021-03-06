package stringProblem

func reverseWordsiii(s string) string {
	if len(s) == 0 {
		return ""
	}

	ret := make([]byte, 0)
	length := len(s)
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}

		for p := start; p < i; p++ {
			ret = append(ret, s[start+i-1-p])
		}

		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
