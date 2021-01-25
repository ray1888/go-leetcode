package tree

import (
	"fmt"
	"math"
)

func verifyPostorderRec(postorder []int) bool {
	return rec(0, len(postorder)-1, postorder)
}

func rec(i, j int, postorder []int) bool {
	if i >= j {
		return true
	}
	p := i
	// 查找对于每个根节点的第一个大于它的数，则为右树的长度
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}
	// p ==j 是用于判断右数的数
	return p == j && rec(i, m-1, postorder) && rec(m, j-1, postorder)
}

func verifyPostorder(postorder []int) bool {
	stack := make([]int, 0)
	root := math.MaxInt32
	for i := len(postorder) - 1; i >= 0; i-- {
		fmt.Printf("postOrder i is %d, root is %d\n", postorder[i], root)
		if postorder[i] > root {
			return false
		}
		/*
		 右树部分是不会进入的分支的，因为自从压进去了右的根节点后
		 ，postorder[i] > 右数的根节点（从二叉搜索树定义可以得知）
		 因为如果是搜索二叉树，遍历到右树完成后，root的值才会改变
		*/
		if len(stack) > 0 && postorder[i] < stack[len(stack)-1] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
