package backtracing

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	result := make([][]int, 0)
	// used:= make([]bool, len(candidates))
	tmp := make([]int, 0)
	combineDfsI(candidates, 0, target, tmp, &result)
	return result
}

func combineDfsI(candidates []int, start, target int, tmp []int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		data := make([]int, len(tmp))
		copy(data, tmp)
		*result = append(*result, data)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		tmp = append(tmp, candidates[i])
		combineDfsI(candidates, i, target-candidates[i], tmp, result)
		tmp = tmp[:len(tmp)-1]
	}
}
