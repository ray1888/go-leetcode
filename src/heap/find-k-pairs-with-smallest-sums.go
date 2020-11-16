package heap

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return [][]int{}
	}
	n := len(nums1)
	m := len(nums2)
	pq := MatrixHeap{}
	heap.Init(&pq)
	for i := 0; i < n && i < k; i++ {
		heap.Push(&pq, MatrixItem{Col: i, Row: 0, Val: nums1[i] + nums2[0]})
	}

	result := make([][]int, 0)

	for i := 0; i < k && len(pq) != 0; i++ {
		minElem := heap.Pop(&pq).(MatrixItem)
		result = append(result, []int{nums1[minElem.Col], nums2[minElem.Row]})
		minElem.Row++
		if minElem.Row < m {
			minElem.Val = nums1[minElem.Col] + nums2[minElem.Row]
			heap.Push(&pq, minElem)
		}
	}
	return result
}
