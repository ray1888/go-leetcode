package minstack

type MinStack struct {
	List         []int
	MinValueList []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{}
	m.List = make([]int, 0)
	m.MinValueList = make([]int, 0)
	return m
}

func (this *MinStack) Push(x int) {
	if len(this.MinValueList) == 0 {
		this.MinValueList = append(this.MinValueList, x)
	} else {
		if x < this.MinValueList[0] {
			this.MinValueList = append(this.MinValueList, 0)
			copy(this.MinValueList[1:], this.MinValueList[:])
			this.MinValueList[0] = x
		} else {
			k := this.MinValueList[0]
			this.MinValueList = append(this.MinValueList, 0)
			copy(this.MinValueList[1:], this.MinValueList[:])
			this.MinValueList[0] = k
		}
	}
	// fmt.Println(this.MinValueList)
	this.List = append(this.List, 0)
	copy(this.List[1:], this.List[:])
	this.List[0] = x
}

func (this *MinStack) Pop() {
	if len(this.List) >= 1 {
		this.List = this.List[1:]
		this.MinValueList = this.MinValueList[1:]
	}
}

func (this *MinStack) Top() int {
	return this.List[0]
}

func (this *MinStack) GetMin() int {
	// fmt.Println(this.MinValueList)
	return this.MinValueList[0]
}
