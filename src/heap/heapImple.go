package heap

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TupleHeap [][]int

func (th TupleHeap) Len() int { return len(th) }
func (th TupleHeap) Less(i, j int) bool {
	return th[i][1] > th[j][1]
}
func (th TupleHeap) Swap(i, j int) { th[i], th[j] = th[j], th[i] }

func (th *TupleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*th = append(*th, x.([]int))
}

func (th *TupleHeap) Pop() interface{} {
	old := *th
	n := len(old)
	x := old[n-1]
	*th = old[0 : n-1]
	return x
}

type MatrixItem struct {
	Col int
	Row int
	Val int
}

type MatrixHeap []MatrixItem

func (h MatrixHeap) Len() int           { return len(h) }
func (h MatrixHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MatrixHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MatrixHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(MatrixItem))
}

func (h *MatrixHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
