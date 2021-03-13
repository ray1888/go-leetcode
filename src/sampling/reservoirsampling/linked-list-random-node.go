package reservoirsampling

import (
	"go-leetcode/src/datastructure"
	"math/rand"
	"time"
)

type Solution struct {
	r    *rand.Rand
	head *datastructure.ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *datastructure.ListNode) Solution {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	solution := Solution{
		head: head,
		r:    r,
	}
	return solution
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	i := 2
	cur := this.head.Next
	val := this.head.Val
	for cur != nil {
		if this.r.Intn(i)+1 == i {
			val = cur.Val
		}
		i++
		cur = cur.Next
	}
	return val
}
