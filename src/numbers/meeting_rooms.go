package numbers

import "sort"

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

func MeettingRooms(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	r := make(roomIntervals, 0)
	for _, item := range intervals {
		newItem := roomInterval{}
		newItem.Start = item[0]
		newItem.End = item[1]
		r = append(r, newItem)
	}

	sort.Sort(&r)

	for i := 0; i < len(r)-1; i++ {
		if roomIntervals[i].End > roomInterval[i+1].Start {
			return false
		}
	}
	return true
}
