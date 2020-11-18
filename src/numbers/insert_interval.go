package numbers

func insert(intervals [][]int, newInterval []int) [][]int {
	tmp := make([][]int, len(intervals)+1)
	size := 0
	p := 0

	for ; p < len(intervals) && intervals[p][0] < newInterval[0]; p++ {
		tmp[size] = intervals[p]
		size++
	}

	if size == 0 || tmp[size-1][1] < newInterval[0] {
		tmp[size] = newInterval
		size++
	} else {
		tmp[size-1][1] = max(newInterval[1], tmp[size-1][1])
	}

	for ; p < len(intervals); p++ {
		if tmp[size-1][1] < intervals[p][0] {
			tmp[size] = intervals[p]
			size++
		} else {
			tmp[size-1][1] = max(tmp[size-1][1], intervals[p][1])
		}
	}

	return tmp[:size]
}
