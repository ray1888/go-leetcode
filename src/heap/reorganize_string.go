package heap

import (
	"container/heap"
	"strings"
)

func reorganizeString(S string) string {
	if len(S) == 0 {
		return ""
	}
	freq := make([]int, 26)
	for _, item := range S {
		i := int(item - 'a')
		freq[i]++
	}

	maxFreq := 0
	h := TupleHeap{}
	heap.Init(&h)
	for index, item := range freq {
		if item == 0 {
			continue
		}
		if item > maxFreq {
			maxFreq = item
		}
		newItem := []int{index, item}
		heap.Push(&h, newItem)
	}
	if maxFreq > (len(S) + 1/2) {
		return ""
	}

	result := make([]string, 0)
	for h.Len() > 1 {
		first := heap.Pop(&h).([]int)
		second := heap.Pop(&h).([]int)
		result = append(result, string(rune(first[0])+'a'))
		result = append(result, string(rune(second[0])+'a'))
		first[1]--
		second[1]--
		if first[1] > 0 {
			heap.Push(&h, first)
		}
		if second[1] > 0 {
			heap.Push(&h, second)
		}
	}
	if h.Len() == 0 {
		return strings.Join(result, "")
	}
	peek := heap.Pop(&h).([]int)
	if peek[1] != 1 {
		return ""
	}
	result = append(result, string(rune(peek[0])+'a'))
	return strings.Join(result, "")
}
