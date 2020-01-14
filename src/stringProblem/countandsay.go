package stringProblem

import (
	"strconv"
	"strings"
)

func CountAndSay(n int) string {
	if n < 1 {
		return ""
	}
	s := "1"
	for i := 1; i < n; i++ {
		count := 1
		sb := make([]string, 0)
		for p := 0; p < len(s); p++ {
			if p+1 < len(s) && s[p] == s[p+1] {
				count += 1
			} else {
				sb = append(sb, strconv.Itoa(count))
				sb = append(sb, string(s[p]))
				count = 1
			}
		}
		s = strings.Join(sb, "")
	}
	return s
}
