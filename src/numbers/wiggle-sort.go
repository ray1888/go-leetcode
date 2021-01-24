package numbers

import "sort"

func wiggleSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	sort.Ints(nums)
	for i := 1; i < len(nums)-1; i += 2 {
		swap(nums, i, i+1)
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
