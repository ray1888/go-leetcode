package circlequeue

import "go-leetcode/src/datastructure"

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyCircularQueue struct {
	Capacity int
	Count    int
	Head     *datastructure.ListNode
	Tail     *datastructure.ListNode
}

func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{}
	cq.Capacity = k
	return cq
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Capacity == this.Count {
		return false
	}
	if this.Count == 0 {
		newNode := datastructure.ListNode{Val: value}
		this.Head = &newNode
		this.Tail = this.Head
	} else {
		newNode := datastructure.ListNode{Val: value}
		this.Tail.Next = &newNode
		this.Tail = this.Tail.Next
	}
	this.Count++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Count == 0 {
		return false
	}
	this.Head = this.Head.Next
	this.Count--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.Count == 0 {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.Count == 0 {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Count == this.Capacity
}
