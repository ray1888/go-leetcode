package DynamicProgramming

import "math"

/*
	粉刷房子二：
	用k个颜色粉刷房子，是原来medium的粉刷房子的变化版本，把3变成了K。

	思路与原来的题目类似，但是偏码技巧上，因为他是需要对k个颜色同时进行对比，
	因此如果直接用一个一维数组去更新的话，会把部分历史数据更新的话，回导致每次取值的时候，最小值不一定正确
	因此此处使用二维数组来进行处理

	压缩空间的思路：
	1. 可以每次进行的时候先把上次的历史先Copy一份，然后再对copy继续去除c=j之后的排序操作
*/

func MinCostWithKColors(nums [][]int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	k := len(nums[0])

	d := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		d[i] = make([]int, k)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < k; j++ {
			minV := math.MaxInt32
			for c := 0; c < k; c++ {
				if c != j {
					minV = min(minV, d[i-1][c])
				}
			}
			d[i][j] = minV + nums[i-1][j]
		}
	}
	minValue := math.MaxInt32
	for i := 0; i < k; i++ {
		minValue = min(minValue, d[n][i])
	}
	return minValue
}
