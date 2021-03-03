package stringProblem

import "strconv"

func compareVersion(version1 string, version2 string) int {
	i := 0
	j := 0

	for i < len(version1) || j < len(version2) {
		var tm1 = ""
		var tm2 = ""
		for i < len(version1) && version1[i] != '.' {
			tm1 += string(version1[i])
			i++
		}
		for j < len(version2) && version2[j] != '.' {
			tm2 += string(version2[j])
			j++
		}
		in1, _ := strconv.Atoi(tm1)
		in2, _ := strconv.Atoi(tm2)
		if in1 > in2 {
			return 1
		} else if in1 < in2 {
			return -1
		}
		i++
		j++

	}
	return 0
}
