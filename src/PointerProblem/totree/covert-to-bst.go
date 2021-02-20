package totree

import "go-leetcode/src/datastructure"

func sortedListToBST(head *datastructure.ListNode) *datastructure.TreeNode {
	return buildTree(head, nil)
}

func getMedian(left, right *datastructure.ListNode) *datastructure.ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right *datastructure.ListNode) *datastructure.TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &datastructure.TreeNode{Val: mid.Val, Left: nil, Right: nil}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}
