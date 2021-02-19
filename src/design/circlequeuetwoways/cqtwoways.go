package circlequeuetwoways

type ListNode struct {
	Val  int
	Next *ListNode
	Pre  *ListNode
}

type MyCircularDeque struct {
	Head     *ListNode
	Tail     *ListNode
	Count    int
	Capacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	cd := MyCircularDeque{}
	cd.Capacity = k
	return cd
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Capacity == this.Count {
		return false
	}
	newNode := ListNode{Val: value}
	if this.Count == 0 {
		this.Head = &newNode
		this.Tail = &newNode
	} else {
		this.Head.Pre = &newNode
		newNode.Next = this.Head
		this.Head = this.Head.Pre
	}
	this.Count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Capacity == this.Count {
		return false
	}
	newNode := ListNode{Val: value}
	if this.Count == 0 {
		this.Head = &newNode
		this.Tail = &newNode
	} else {
		this.Tail.Next = &newNode
		newNode.Pre = this.Tail
		this.Tail = this.Tail.Next
	}
	this.Count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.Count == 0 {
		return false
	} else if this.Count == 1 {
		this.Head = nil
		this.Tail = nil
	} else {
		newHead := this.Head.Next
		this.Head.Next = nil
		this.Head = newHead
		this.Head.Pre = nil
	}
	this.Count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.Count == 0 {
		return false
	} else if this.Count == 1 {
		this.Head = nil
		this.Tail = nil
	} else {
		newTail := this.Tail.Pre
		this.Tail.Pre = nil
		this.Tail = newTail
		this.Tail.Next = nil
	}
	this.Count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.Count == 0 {
		return -1
	}
	return this.Head.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.Count == 0 {
		return -1
	}
	return this.Tail.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Count == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Capacity == this.Count
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
