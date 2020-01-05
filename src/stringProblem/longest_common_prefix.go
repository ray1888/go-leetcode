package stringProblem

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	head := strs[0]
	for i := 0; i < len(head); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != head[i] {
				return head[0:i]
			}
		}
	}
	return head
}
