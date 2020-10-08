package backtracing

func dfsPhone(digits string, index int, elem []byte, dMap map[byte][]byte, result *[]string) {
	if index == len(digits) {
		newElem := make([]byte, len(elem))
		copy(newElem, elem)
		*result = append(*result, string(newElem))
		return
	}
	val, ok := dMap[digits[index]]
	if !ok {
		return
	}
	for _, item := range val {
		elem = append(elem, item)
		dfsPhone(digits, index+1, elem, dMap, result)
		elem = elem[:len(elem)-1]
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digitMap := make(map[byte][]byte)
	digitMap['2'] = []byte{'a', 'b', 'c'}
	digitMap['3'] = []byte{'d', 'e', 'f'}
	digitMap['4'] = []byte{'g', 'h', 'i'}
	digitMap['5'] = []byte{'j', 'k', 'l'}
	digitMap['6'] = []byte{'m', 'n', 'o'}
	digitMap['7'] = []byte{'p', 'q', 'r', 's'}
	digitMap['8'] = []byte{'t', 'u', 'v'}
	digitMap['9'] = []byte{'w', 'x', 'y', 'z'}
	result := make([]string, 0)
	elem := make([]byte, 0)
	dfsPhone(digits, 0, elem, digitMap, &result)
	return result
}
