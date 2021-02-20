package binary_search

/*
	寻找重复数字出现的个数
	给定数组[1,2,2,2,5,6],重复数字有且只有一个
	output: 3

*/

func bSearch(arr []int, target int, isLeft bool) int {
	pos := 0
	i := 0
	j := len(arr) - 1
	for i <= j {
		mid := i + (j-i)/2
		if arr[mid] > target {
			j = mid - 1
		} else if arr[mid] < target {
			i = mid + 1
		} else {
			pos = mid
			if isLeft {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}
	return pos
}

func FindRepeatCountOLogN(arr []int, repeatNum int) int {
	leftIndex := bSearch(arr, repeatNum, true)
	rightIndex := bSearch(arr, repeatNum, false)
	return rightIndex - leftIndex + 1
}
