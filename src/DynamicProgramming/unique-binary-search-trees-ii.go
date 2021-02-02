package DynamicProgramming

import (
	"go-leetcode/src/datastructure"
)

func generateTrees(n int) []*datastructure.TreeNode {
	if n <= 0 {
		return []*datastructure.TreeNode{}
	}
	return generateTreesSE(1, n)
}

func generateTreesSE(start, end int) []*datastructure.TreeNode {
	if start > end {
		return nil
	}
	if end == 1 || start == end {
		return []*datastructure.TreeNode{{ // returns list with a node if only one value
			Val: start,
		}}
	}
	if end == 2 {
		return []*datastructure.TreeNode{
			{
				Val:   start,
				Right: &datastructure.TreeNode{Val: start + 1},
			},
			{
				Val:  start + 1,
				Left: &datastructure.TreeNode{Val: start},
			},
		}
	}

	list := make([]*datastructure.TreeNode, 0, end-start)
	for i := start; i <= end; i++ {
		left := generateTreesSE(start, i-1) // get the left nodes
		right := generateTreesSE(i+1, end)  // get the right nodes
		if left == nil {
			for _, r := range right {
				list = append(list, &datastructure.TreeNode{
					Val:   i,
					Right: r,
				})
			}
			continue
		}
		if right == nil {
			for _, l := range left {
				list = append(list, &datastructure.TreeNode{
					Val:  i,
					Left: l,
				})
			}
			continue
		}
		for _, l := range left {
			for _, r := range right {
				list = append(list, &datastructure.TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return list
}
func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return -1
}
