package tree

func levelOrderNAryTree(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)

	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		tmpResult := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[i]
			tmpResult = append(tmpResult, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
		queue = queue[length:]
		result = append(result, tmpResult)
	}
	return result
}
