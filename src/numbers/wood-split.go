package numbers

import "go-leetcode/src/utils"

/*
题目描述
给定长度为n的数组，每个元素代表一个木头的长度，木头可以任意截断，从这堆木头中截出至少k个相同长度为m的木块。已知k，求max(m)。

ps: 截断的长度必须是整数

输入两行，第一行n, k，第二行为数组序列。输出最大值。

输入
5 5
4 7 2 10 5
输出
4

*/

func splitBrutalForce(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}
	maxValue := 0
	for _, item := range arr {
		maxValue = utils.Max(item, maxValue)
	}
	m := 1
	maxCount := 0
	for m < maxValue {
		count := 0
		for i := 0; i < len(arr); i++ {
			count += arr[i] / m
		}
		if count >= k {
			maxCount = utils.Max(maxCount, m)
		}
		m++
	}
	return maxCount
}

func splitHalfCut(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}
	r := -1
	l := 1
	for _, item := range arr {
		r = utils.Max(item, r)
	}

	for l < r {
		mid := l + (r-l)/2
		f := func(mid int) int {
			res := 0
			for i := 0; i < len(arr); i++ {
				res += arr[i] / mid
			}
			return res
		}
		res := f(mid)
		if res >= k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
