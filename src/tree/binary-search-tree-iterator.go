package tree

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.stack = make([]*TreeNode, 0)
	it.PushLeft(root)
	return it
}

func (this *BSTIterator) PushLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}

}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	item := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.PushLeft(item.Right)
	return item.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.stack) == 0 {
		return false
	}
	return true
}
