package tree

import "go-leetcode/src/datastructure"

//func _lowestCommonAncestorBT(root, p, q *datastructure.TreeNode, result *[]*datastructure.TreeNode,
//	record *map[int][]*datastructure.TreeNode) {
//	if root == nil {
//		return
//	}
//	if root.Val == p.Val || root.Val == q.Val {
//		*result = append(*(result), root)
//		var ssr *[]*datastructure.TreeNode
//		copy(*ssr, *result)
//		if res, ok := *record[root.Val]; !ok {
//			*record[root.Val] = make([]*datastructure.TreeNode, 0)
//		}
//		*record[root.Val] = ssr
//	}
//	*result = append(*(result), root)
//	_lowestCommonAncestorBT(root.Left, p, q, result, record)
//	_lowestCommonAncestorBT(root.Right, p, q, result, record)
//}
//
//func lowestCommonAncestorBT(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
//	if root == nil {
//		return nil
//	}
//	result := make(map[int][]*datastructure.TreeNode)
//	sr := make([]*datastructure.TreeNode, 0)
//	_lowestCommonAncestorBT(root, p, q, &sr, &result)
//	p_map := result[p.Val]
//	q_map := result[q.Val]
//	index := -1
//	for i := 0; i < min(len(q_map), len(p_map)); i++ {
//		if p_map[i] == q_map[i] {
//			index = 1
//		} else {
//			break
//		}
//	}
//	if index == -1 {
//		return nil
//	} else {
//		return p_map[index]
//	}
//}

func lowestCommonAncestorBT(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestorBT(root.Left, p, q)
	right := lowestCommonAncestorBT(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestorDFS(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
	parent := map[int]*datastructure.TreeNode{}
	visited := map[int]bool{}

	var dfs func(*datastructure.TreeNode)
	dfs = func(r *datastructure.TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}
