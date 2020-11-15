package heap

import "container/heap"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix[0])
	m := len(matrix)
	pq := MatrixHeap{}
	heap.Init(&pq)
	for i := 0; i < m && i < k; i++ {
		heap.Push(&pq, MatrixItem{Col: 0, Row: i, Val: matrix[i][0]})
	}

	for i := 0; i < k-1; i++ {
		minElem := heap.Pop(&pq).(MatrixItem)
		minElem.Col++
		if minElem.Col < n {
			minElem.Val = matrix[minElem.Row][minElem.Col]
			heap.Push(&pq, minElem)
		}
	}
	return pq[0].Val
}
