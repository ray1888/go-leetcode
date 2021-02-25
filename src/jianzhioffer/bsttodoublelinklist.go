package jianzhioffer

import "go-leetcode/src/datastructure"

func bstTODoubleLinkList(root *datastructure.TreeNode) *datastructure.DoubleLinkListNode {
	result := make([]int, 0)
	midTravel(root, &result)
	head := makeDoubleLinkList(result)
	return head
}

func midTravel(node *datastructure.TreeNode, result *[]int) {
	if node == nil {
		return
	}
	midTravel(node.Left, result)
	*result = append(*result, node.Val)
	midTravel(node.Right, result)
}

func makeDoubleLinkList(arr []int) *datastructure.DoubleLinkListNode {
	var head, tail *datastructure.DoubleLinkListNode
	head.Next = tail
	head.Pre = tail
	tail.Pre = head
	tail.Next = head
	work := head
	for i := 0; i < len(arr); i++ {
		newNode := new(datastructure.DoubleLinkListNode)
		newNode.Pre = work
		newNode.Next = tail
		tail.Pre = newNode
		work.Next = newNode
		work = work.Next
	}
	return head
}
