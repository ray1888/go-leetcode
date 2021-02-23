package numbers

import "sort"

var result [][]int

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	sort.Ints(nums)
	result = make([][]int, 0)
	tmp := make([]int, 0)
	dfs(nums, tmp, 0)
	return result
}

func dfs(nums, tmp []int, start int) {
	tmp2 := make([]int, len(tmp))
	copy(tmp2, tmp)
	result = append(result, tmp2)
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		dfs(nums, tmp, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
