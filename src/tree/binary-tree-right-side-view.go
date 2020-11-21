package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		length := len(q)
		result = append(result, q[0].Val)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
		}
	}
	return result
}
