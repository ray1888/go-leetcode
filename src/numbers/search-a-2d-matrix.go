package numbers

/*
	题解：
	这道题本质上是把一个排好序的二维数组可以理解为用二分查找来缩小范围
	因为默认排序已经建好了，所以只要把开始和结束下标放到第一个元素和矩阵的最后一个元素
	即可进行二分查找
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	i := 0
	row := len(matrix)
	col := len(matrix[0])
	j := row*col - 1

	for i <= j {
		mid := i + (j-i)/2
		midCol := mid % col
		midRow := mid / col
		if target == matrix[midRow][midCol] {
			return true
		}
		if target < matrix[midRow][midCol] {
			j = mid - 1
		}
		if target > matrix[midRow][midCol] {
			i = mid + 1
		}
	}
	return false
}
