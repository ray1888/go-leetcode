package stock

import (
	"go-leetcode/src/utils"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minVal := int(math.Pow(2, 32))
	cur := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minVal {
			minVal = prices[i]
		}
		cur = utils.Max(cur, prices[i]-minVal)
	}
	return cur
}
