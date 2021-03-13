package slidingwindow

import "go-leetcode/src/utils"

func characterReplacement(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}
	left := 0
	right := 0
	charArray := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		charArray[i] = s[i]
	}

	res := 0
	maxCount := 0
	freq := make([]int, 26)
	for right < len(s) {
		freq[charArray[right]-'A']++
		maxCount = utils.Max(maxCount, freq[charArray[right]-'A'])
		right++
		// 说明此时 k 不够用
		// 把其它不是最多出现的字符替换以后，都不能填满这个滑动的窗口，这个时候须要考虑左边界向右移动
		// 移出滑动窗口的时候，频数数组须要相应地做减法
		if right-left > maxCount+k {
			freq[charArray[left]-'A']--
			left++
		}
		res = utils.Max(res, right-left)
	}
	return res
}
