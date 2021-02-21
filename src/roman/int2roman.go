package roman

import "strings"

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	alpha := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var sb strings.Builder

	for i := 0; i < len(nums) && num >= 0; i++ {
		for nums[i] <= num {
			num -= nums[i]
			sb.WriteString(alpha[i])
		}
	}

	return sb.String()
}
