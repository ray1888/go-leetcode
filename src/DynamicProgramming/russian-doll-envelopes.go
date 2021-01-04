package DynamicProgramming

import (
	"fmt"
	"sort"
)

type IntArray [][]int

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Swap(i, j int) {
	ia[i], ia[j] = ia[j], ia[i]
}

func (ia IntArray) Less(i, j int) bool {
	if ia[i][0] == ia[j][0] {
		return ia[i][1] > ia[j][1]
	}
	return ia[i][0] < ia[j][0]
}

func _BinarySearchInsPos(d []int, end int, target int) int {
	start := 0
	end = end - 1
	for start <= end {
		mid := start + (end-start)/2
		if d[mid] > target {
			end = mid - 1
		} else if d[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	d := make([]int, len(nums))
	lengthLIS := 0
	for _, val := range nums {
		index := _BinarySearchInsPos(d, lengthLIS, val)
		d[index] = val
		if index == lengthLIS {
			lengthLIS += 1
		}
	}
	return lengthLIS
}

func maxEnvelopes(envelopes [][]int) int {
	en := IntArray(envelopes)
	sort.Sort(en)
	height := make([]int, 0, len(en))
	fmt.Println(en)
	for _, item := range en {
		height = append(height, item[1])
	}
	return lengthOfLIS(height)
}
