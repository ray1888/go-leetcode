package stringProblem

import "strings"

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	var res string
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		//cur := strings.TrimSpace(arr[i]) 更加严谨的做法应该还要去掉末尾的空格
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 {
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return "/" + res
}
