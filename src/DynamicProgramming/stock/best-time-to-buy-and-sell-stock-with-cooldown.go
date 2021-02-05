package stock

import (
	"go-leetcode/src/utils"
	"math"
)

func maxProfitCoolDown(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dpI0 := 0
	dpI1 := math.MinInt32
	dpPre0 := 0

	for i := 0; i < len(prices); i++ {
		tmp := dpI0
		// 第n-1天已经不持股和第n-1天持股并且再今天卖出的价钱
		dpI0 = utils.Max(dpI0, dpI1+prices[i])
		// 第n-1天已经持股和第n-2天不持股并且今天买入的价钱
		dpI1 = utils.Max(dpI1, dpPre0-prices[i])
		dpPre0 = tmp
	}
	return dpI0
}
