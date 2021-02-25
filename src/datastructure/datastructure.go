package datastructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

type ConnectNode struct {
	Val   int
	Left  *ConnectNode
	Right *ConnectNode
	Next  *ConnectNode
}

type DoubleLinkListNode struct {
	Val  int
	Next *DoubleLinkListNode
	Pre  *DoubleLinkListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenListNode(data []int) *ListNode {
	head := new(ListNode)
	p := head
	for _, value := range data {
		p.Val = value
		p.Next = new(ListNode)
		p = p.Next
	}
	return head.Next
}
