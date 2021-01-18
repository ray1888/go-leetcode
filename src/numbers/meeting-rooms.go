package numbers

import "sort"

type Intervals [][]int

func (I Intervals) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func (I Intervals) Less(i, j int) bool {
	return I[i][0] < I[j][0]
}

func (I Intervals) Len() int {
	return len(I)
}

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}
	k := Intervals(intervals)
	sort.Sort(k)
	lastEnd := k[0][1]
	for i := 1; i < len(k); i++ {
		if lastEnd > k[i][0] {
			return false
		}
		lastEnd = k[i][1]
	}
	return true
}
