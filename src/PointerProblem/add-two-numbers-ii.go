package PointerProblem

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	stackA := make([]int, 0)
	stackB := make([]int, 0)

	for l1 != nil {
		stackA = append(stackA, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stackB = append(stackB, l2.Val)
		l2 = l2.Next
	}

	dummy := new(ListNode)
	head := dummy
	pos := 0
	result := make([]int, 0)
	for len(stackA) > 0 || len(stackB) > 0 || pos > 0 {
		sum := 0
		if len(stackA) > 0 {
			sum += stackA[len(stackA)-1]
			stackA = stackA[:len(stackA)-1]
		}
		if len(stackB) > 0 {
			sum += stackB[len(stackB)-1]
			stackB = stackB[:len(stackB)-1]
		}
		sum += pos
		pos = sum / 10
		sum %= 10
		result = append(result, sum)
	}
	for i := len(result) - 1; i >= 0; i-- {
		newNode := new(ListNode)
		head.Next = newNode
		newNode.Val = result[i]
		head = head.Next
	}
	head.Next = nil
	return dummy.Next
}
