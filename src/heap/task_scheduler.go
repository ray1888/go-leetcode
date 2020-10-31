package heap

import (
	"container/heap"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// An IntHeap is a min-heap of ints.

func leastInterval(tasks []byte, n int) int {
	charSlice := [26]int{}

	for i := 0; i < len(tasks); i++ {
		charSlice[tasks[i]-'A'] += 1
	}
	pq := IntHeap{}
	heap.Init(&pq)
	for _, item := range charSlice {
		if item != 0 {
			heap.Push(&pq, item)
		}
	}
	result := 0
	idle := 0
	for len(pq) > 0 {
		result += (n + 1)
		idle = n + 1 - len(pq)
		size := min(n+1, len(pq))
		for i := 0; i < size; i++ {
			charSlice[i] = heap.Pop(&pq).(int) - 1
		}
		for i := 0; i < size; i++ {
			if charSlice[i] != 0 {
				heap.Push(&pq, charSlice[i])
			}
		}
	}
	return result - idle
}
