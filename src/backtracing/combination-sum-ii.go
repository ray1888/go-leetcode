package backtracing

import (
	"fmt"
	"sort"
)

// TODO 解决去重问题
func combinationSum2NotUnique(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	result := make([][]int, 0)
	fmt.Println(candidates)
	for i := 0; i < len(candidates); i++ {
		tmp := make([]int, 0)
		combineDfs(candidates, i, target, tmp, &result)
	}
	return result
}

func combineDfs(candidates []int, start, target int, tmp []int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		if start == 4 {
			fmt.Println(start, tmp)
		}
		data := make([]int, len(tmp))
		copy(data, tmp)
		*result = append(*result, data)
		return
	}

	for i := start; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		combineDfs(candidates, i+1, target-candidates[i], tmp, result)
		tmp = tmp[:len(tmp)-1]
	}
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var sequence []int
	var dfs func(pos, rest int)
	dfs = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), sequence...))
			return
		}
		if pos == len(freq) || rest < freq[pos][0] {
			return
		}

		dfs(pos+1, rest)

		most := min(rest/freq[pos][0], freq[pos][1])
		for i := 1; i <= most; i++ {
			sequence = append(sequence, freq[pos][0])
			dfs(pos+1, rest-i*freq[pos][0])
		}
		sequence = sequence[:len(sequence)-most]
	}
	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
