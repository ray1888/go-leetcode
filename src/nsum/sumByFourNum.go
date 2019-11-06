package nsum

import "sort"

func NSumBy4Nums(input []int, target int) [][]int {
	if len(input) == 0 {
		return [][]int{}
	}
	n := len(input)
	result := make([][]int, 0, 10)
	sort.Ints(input)
	for p := n - 1; p >= 3; p-- {
		if 4*input[p] < target {
			break
		}
		for k := p - 1; k >= 2; k-- {
			if 3*input[k]+input[p] < target {
				break
			}
			newTarget := (target - input[p] - input[k])
			i := 0
			j := k - 1
			for i < j {
				if input[i]+input[j] == newTarget {
					result = append(result, []int{input[i], input[j], input[k], input[p]})
					for i < j && input[i+1] == input[i] {
						i += 1
					}
					for i < j && input[j] == input[j-1] {
						j -= 1
					}
					i += 1
					j -= 1
				} else if input[i]+input[j] > newTarget {
					j -= 1
				} else {
					i += 1
				}
			}
			for k >= 2 && input[k-1] == input[k] {
				k -= 1
			}
		}
		for p >= 2 && input[p-1] == input[p] {
			p -= 1
		}
	}
	return result
}
