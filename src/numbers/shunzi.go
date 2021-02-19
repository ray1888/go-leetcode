package numbers

import "sort"

/*
	从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。


	从题目可以获取到如下信息：
	1. 0的出现次数最多为两次（大王和小王），其他只能出现一次
	2. 最大和最小值的差不能大于等于5


	所以0,0,1,2,5 = 》 可以理解为1,2，3，4,5


*/

func isStraight(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}
