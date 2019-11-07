package graph

var mapping []string = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func combination(digits string, idx int, str string, result *[]string) {
	if idx == len(digits) {
		*result = append(*result, str)
		return
	}
	char := mapping[digits[idx]-'2']
	for i := 0; i < len(char); i++ {
		combination(digits, idx+1, str+string(char[i]), result)
	}
}

func letterCombinationsRecursive(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0, 10)
	combination(digits, 0, "", &result)
	return result
}
