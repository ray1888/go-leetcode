package numbers

import "sort"

var result2 [][]int

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	sort.Ints(nums)
	result = make([][]int, 0)
	tmp := make([]int, 0)
	dfssubset2(nums, tmp, 0)
	return result2
}

func dfssubset2(nums, tmp []int, start int) {
	tmp2 := make([]int, len(tmp))
	copy(tmp2, tmp)
	result2 = append(result2, tmp2)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		dfssubset2(nums, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
