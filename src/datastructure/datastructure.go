package datastructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ConnectNode struct {
	Val   int
	Left  *ConnectNode
	Right *ConnectNode
	Next  *ConnectNode
}
