package binary_search

/*

	对于回旋数组这种问题，一般的处理方法是
	二分法

	回旋数组
	[....,0,....]
	递增区间， 0， 递增区间
	但是本题目比较特殊，因为它还要判断target值是否在数组中，则应该这样处理
	有四种情况
	1.判断目标值是否在中间数小，并且在左边的递增区间 => j往左收
	2.判断目标值是否在中间数大，或者在不在左边的递增区间 => i往右移动
	3.判断目标值是否在中间数大，并且在右边的递增区间 => i往右移动
	4.判断目标值是否在中间数小，或者不在右边的递增区间 =>  j往左收

*/

func searchRotateArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return mid
		}
		if nums[mid] >= nums[low] {
			if target < nums[mid] && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
