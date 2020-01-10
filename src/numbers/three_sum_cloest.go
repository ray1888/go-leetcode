package numbers

import "sort"

func ThreeSumCloest(input []int, target int) int {
	if len(input) == 0 {
		return 0
	}
	sort.Ints(input)
	result := 0
	minDiff := (2 << 30)
	for k := len(input) - 1; k >= 2; k-- {
		i := 0
		j := k - 1
		for i < j {
			sum := input[i] + input[j] + input[k]
			if sum == target {
				return sum
			} else if sum < target {
				i += 1
			} else {
				j -= 1
			}
			diff := target - sum
			if diff < 0 {
				diff = -1 * diff
			}
			if diff < minDiff {
				diff = -1 * diff
			}

		}
	}
	return result
}
