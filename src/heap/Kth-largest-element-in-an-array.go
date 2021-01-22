package heap

import "container/heap"

func swap(nums []int, a, b int) {
	tmp := nums[b]
	nums[b] = nums[a]
	nums[a] = tmp
}

func partition(nums []int, start, end int) int {
	i := start
	j := end
	pivot := nums[end]
	for i < j {
		for (nums[i] >= pivot) && (i < j) {
			i += 1
		}
		if i < j {
			swap(nums, i, j)
		}

		for (nums[j] < pivot) && (i < j) {
			j -= 1
		}
		if i < j {
			swap(nums, i, j)
		}

	}
	return i
}

func findKthLargestPartition(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return -1
	}
	low := 0
	high := len(nums) - 1
	for low <= high {
		p := partition(nums, low, high)
		if p == k-1 {
			return nums[p]
		} else if p > k-1 {
			high = p - 1
		} else {
			low = p + 1
		}
	}
	return -1
}

func findKthLargestMinHeap(nums []int, k int) int {
	pq := IntHeap{}
	heap.Init(&pq)
	for i := 0; i < len(nums); i++ {
		if pq.Len() >= k {
			if pq[0] < nums[i] {
				_ = heap.Pop(&pq)
				heap.Push(&pq, nums[i])
			}
		} else {
			heap.Push(&pq, nums[i])
		}
	}
	return pq[0]
}
