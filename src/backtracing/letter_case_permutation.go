package backtracing

func isAlpha(char byte) bool {
	if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
		return true
	}
	return false
}

func dfsLetter(S string, index int, result *[]string) {
	if index == len(S) {
		*result = append(*result, S)
		return
	}
	dfsLetter(S, index+1, result)
	if isAlpha(S[index]) {
		s1 := []rune(S)
		s1[index] = s1[index] ^ 32
		ss1 := string(s1)
		dfsLetter(ss1, index+1, result)
	}
}

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	dfsLetter(S, 0, &result)
	return result
}
