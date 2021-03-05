package linkedlist

type MyLinkedList struct {
	size int
	head *Node
}

type Node struct {
	val  int
	next *Node
	pre  *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	head := &Node{}
	head.next = head
	head.pre = head
	return MyLinkedList{
		size: 0,
		head: head,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.size < index+1 {
		return -1
	}
	cur := this.head
	for index >= 0 {
		cur = cur.next
		index--
	}
	return cur.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tmp := &Node{
		val:  val,
		next: this.head.next,
		pre:  this.head,
	}
	this.head.next.pre = tmp
	this.head.next = tmp
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tmp := &Node{
		val:  val,
		next: this.head,
		pre:  this.head.pre,
	}
	this.head.pre.next = tmp
	this.head.pre = tmp
	this.size++
}

/** Add a node of value val before the index-th node in the linked list.
If index equals to the length of linked list, the node will be appended to the end of linked list.
If index is greater than the length, the node will not be inserted.
*/
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.size < index {
		return
	}
	if this.size == index {
		this.AddAtTail(val)
		return
	}
	cur := this.head
	for index > 0 {
		cur = cur.next
		index--
	}
	tmp := &Node{
		val:  val,
		next: cur.next,
		pre:  cur,
	}
	cur.next.pre = tmp
	cur.next = tmp
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.size <= index {
		return
	}
	cur := this.head
	for index >= 0 {
		cur = cur.next
		index--
	}
	cur.pre.next = cur.next
	cur.next.pre = cur.pre
	cur.next = nil
	cur.pre = nil
	this.size--
}

// type Node struct{
//     next *Node
//     prev *Node
//     Val   int
// }

// type MyLinkedList struct {
//     head *Node
//     tail *Node
//     indexs []*Node
// }

// /** Initialize your data structure here. */
// func Constructor() MyLinkedList {
//     mll := MyLinkedList{}
//     mll.head = new(Node)
//     mll.tail = new(Node)
//     mll.head.next = mll.tail
//     mll.tail.prev = mll.head
//     mll.indexs = make([]*Node, 0)
//     return mll
// }

// /** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
// func (this *MyLinkedList) Get(index int) int {
//     if len(this.indexs) < index{
//         return -1
//     }
//     fmt.Println(this.indexs)
//     return this.indexs[index].Val
// }

// /** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
// func (this *MyLinkedList) AddAtHead(val int)  {
//     newnode := new(Node)
//     newnode.Val = val

//     oldNext := this.head.next
//     oldNext.prev = newnode
//     newnode.next = oldNext
//     this.head.next = newnode

//     this.indexs = append([]*Node{newnode}, this.indexs...)
// }

// /** Append a node of value val to the last element of the linked list. */
// func (this *MyLinkedList) AddAtTail(val int)  {
//     newnode := new(Node)
//     newnode.Val = val

//     oldPrev := this.tail.prev
//     oldPrev.next = newnode
//     newnode.prev = oldPrev
//     this.tail.prev = newnode
//     newnode.next = this.tail

//     this.indexs = append(this.indexs, newnode)
// }

// /** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
// func (this *MyLinkedList) AddAtIndex(index int, val int)  {
//     if len(this.indexs) < index{
//         return
//     }

//     oldNode := this.indexs[index]
//     oldPrev := oldNode.prev
//     newNode := new(Node)
//     newNode.Val = val

//     oldNode.prev = newNode
//     newNode.next = oldNode
//     newNode.prev = oldPrev
//     oldPrev.next = newNode

//     preIndex := this.indexs[:index]
//     postIndex := this.indexs[index:]
//     preIndex = append(preIndex, newNode)
//     newItem := append(preIndex, postIndex...)
//     this.indexs =newItem
// }

// /** Delete the index-th node in the linked list, if the index is valid. */
// func (this *MyLinkedList) DeleteAtIndex(index int)  {
//     if len(this.indexs) < index{
//         return
//     }

//     node := this.indexs[index]
//     fmt.Println(node)
//     prev := node.prev
//     next := node.next

//     node.prev = nil
//     node.next = nil
//     prev.next = next
//     next.prev = prev

//     preIndex := this.indexs[:index]
//     postIndex := this.indexs[index+1:]

//     newItem := append(preIndex, postIndex...)
//     this.indexs = newItem
// }
