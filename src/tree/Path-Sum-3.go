package tree

import "go-leetcode/src/datastructure"

func _pathSum3(root *datastructure.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val == sum {
		result += 1
	}
	result += (_pathSum3(root.Left, sum-root.Val) + _pathSum3(root.Right, sum-root.Val))
	return result
}

func pathSum3(root *datastructure.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := _pathSum3(root, sum) + pathSum3(root.Left, sum) + pathSum3(root.Right, sum)
	return result
}

func pathSum3Prefix(root *datastructure.TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := make(map[int]int)
	result[0] = 1
	return pathSumRec(root, sum, 0, result)
}

func pathSumRec(node *datastructure.TreeNode, target, curSum int, sumMap map[int]int) int {
	if node == nil {
		return 0
	}
	res := 0
	curSum += node.Val
	if val, ok := sumMap[curSum-target]; ok {
		res += val
	}
	if val, ok := sumMap[curSum]; ok {
		sumMap[curSum] = val + 1
	} else {
		sumMap[curSum] = 1
	}
	res += pathSumRec(node.Left, target, curSum, sumMap)
	res += pathSumRec(node.Right, target, curSum, sumMap)

	sumMap[curSum]--
	return res
}
