package nonstringlinear

import "go-leetcode/src/utils"

/*

	这题的子问题相当于分解成了，基于上一个状态，求子问题即能被2，3，5的最小值为多少，然后填入数组中
*/

func nthUglyNumber(n int) int {
	uglyNumber := make([]int, n)
	p2 := 0
	p3 := 0
	p5 := 0
	uglyNumber[0] = 1
	for i := 1; i < n; i++ {
		value := utils.Min(utils.Min(uglyNumber[p2]*2, uglyNumber[p3]*3), uglyNumber[p5]*5)
		uglyNumber[i] = value
		if value == uglyNumber[p2]*2 {
			p2++
		}
		if value == uglyNumber[p3]*3 {
			p3++
		}
		if value == uglyNumber[p5]*5 {
			p5++
		}
	}

	return uglyNumber[n-1]
}
