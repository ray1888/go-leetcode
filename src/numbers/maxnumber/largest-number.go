package maxnumber

import (
	"sort"
	"strconv"
)

type strCompare []string

func (s strCompare) Len() int {
	return len(s)
}

func (s strCompare) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s strCompare) Less(i, j int) bool {
	o12 := s[i] + s[j]
	o21 := s[j] + s[i]
	return o12 > o21
}

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	str := make(strCompare, 0)
	for _, item := range nums {
		str = append(str, strconv.Itoa(item))
	}

	sort.Sort(&str)
	if str[0] == "0" {
		return "0"
	}
	var result string
	for _, item := range str {
		result += item
	}

	return result
}
