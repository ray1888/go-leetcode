package tree

import (
	ds "go-leetcode/src/datastructure"
)

func connectRecursive(root *ds.ConnectNode) *ds.ConnectNode {
	if root == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connectRecursive(root.Left)
	connectRecursive(root.Right)
	return root
}

func connectIterative(root *ds.ConnectNode) *ds.ConnectNode {
	if root == nil {
		return root
	}
	leftMost := root
	p := root
	for leftMost.Left != nil {
		p = leftMost
		for p != nil {
			p.Left.Next = p.Right
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
			p = p.Next
		}
		leftMost = leftMost.Left
	}
	return root
}
