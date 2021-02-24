package tree

import (
	"go-leetcode/src/datastructure"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func ConstructorSerialTree() Codec {
	c := Codec{}
	return c
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *datastructure.TreeNode) string {
	if root == nil {
		return "[null]"
	}
	result := make([]string, 0)
	queue := make([]*datastructure.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			if node == nil {
				result = append(result, "null")
			} else {
				result = append(result, strconv.Itoa(node.Val))
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
	}
	counter := 0
	for i := len(result) - 1; i > 0; i-- {
		if result[i] == "null" {
			counter++
		} else {
			break
		}
	}
	result = result[:len(result)-counter]
	rr := "["
	ss := strings.Join(result, ",")
	rrr := rr + ss + "]"
	return rrr
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *datastructure.TreeNode {
	if len(data) == 0 {
		return nil
	}
	data = data[1 : len(data)-1]
	// fmt.Println(data)
	ss := strings.Split(data, ",")
	nodes := make([]*datastructure.TreeNode, len(ss))
	for i := 0; i < len(ss); i++ {
		if ss[i] == "null" {
			nodes[i] = nil
		} else {
			nodes[i] = new(datastructure.TreeNode)
			nodes[i].Val, _ = strconv.Atoi(ss[i])
		}
	}
	i := 1
	for _, item := range nodes {
		if item != nil {
			if i < len(ss) {
				item.Left = nodes[i]
				i++
			}
			if i < len(ss) {
				item.Right = nodes[i]
				i++
			}
		}
	}
	return nodes[0]
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
