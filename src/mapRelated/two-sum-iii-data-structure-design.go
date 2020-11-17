package mapRelated

type TwoSumList struct {
	data  []int
	index map[int]bool
}

func (ts *TwoSumList) Add(a int) {
	for _, item := range ts.data {
		ts.index[item+a] = true
	}
	ts.data = append(ts.data, a)
}

func (ts *TwoSumList) Find(k int) bool {
	_, ok := ts.index[k]
	return ok
}
