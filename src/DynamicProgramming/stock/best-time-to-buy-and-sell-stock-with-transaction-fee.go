package stock

import (
	"go-leetcode/src/utils"
	"math"
)

func maxProfitWithFee(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	dpI0 := 0
	dpI1 := math.MinInt32

	for i := 0; i < len(prices); i++ {
		tmp := dpI0
		// 第n-1天已经不持股和第n-1天持股并且再今天卖出的价钱
		dpI0 = utils.Max(tmp, dpI1+prices[i])
		// 第n-1天已经持股和第n-1天不持股并且今天买入的价钱, 减去手续费
		dpI1 = utils.Max(dpI1, tmp-prices[i]-fee)
		tmp = dpI0
	}
	return dpI0
}
