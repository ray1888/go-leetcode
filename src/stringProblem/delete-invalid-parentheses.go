package stringProblem

import "strings"

func minRemoveToMakeValid(s string) string {
	if len(s) == 0 {
		return s
	}

	stack := make([]int, 0)
	indexList := make([]int, 0)

	for index, item := range s {
		if item == '(' {
			stack = append(stack, index)
		} else if item == ')' {
			if len(stack) == 0 {
				indexList = append(indexList, index)
			} else {
				stack = stack[1:]
			}
		}
	}

	sb := strings.Builder{}

	indexMap := make(map[int]bool)
	for _, item := range indexList {
		indexMap[item] = true
	}
	if len(stack) > 0 {
		for _, index := range stack {
			indexMap[index] = true
		}
	}

	for i := 0; i < len(s); i++ {
		if _, ok := indexMap[i]; ok {
			delete(indexMap, i)
			continue
		}
		sb.WriteByte(s[i])
	}

	return sb.String()
}
