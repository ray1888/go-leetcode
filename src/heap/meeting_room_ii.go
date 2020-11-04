package heap

import (
	"container/heap"
	"sort"
)

type roomInterval struct {
	Start int
	End   int
}

type roomIntervals []roomInterval

func (r roomIntervals) Len() int {
	return len(r)
}

func (r roomIntervals) Less(i, j int) bool {
	return r[i].Start < r[j].Start
}

func (r roomIntervals) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func MeetingRoom(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	r := make(roomIntervals, 0)
	for _, item := range intervals {
		newItem := roomInterval{}
		newItem.Start = item[0]
		newItem.End = item[1]
		r = append(r, newItem)
	}

	sort.Sort(&r)

	rheap := IntHeap{}
	heap.Init(&rheap)
	heap.Push(&rheap, r[0].End)
	for i := 1; i < len(r); i++ {
		if rheap[0] < r[i].Start {
			_ = heap.Pop(&rheap)
		}
		heap.Push(&rheap, r[i].End)
	}
	return rheap.Len()
}
