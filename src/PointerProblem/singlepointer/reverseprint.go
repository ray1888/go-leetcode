package singlepointer

import "go-leetcode/src/datastructure"

func reversePrint(head *datastructure.ListNode) []int {
	if head == nil {
		return []int{}
	}

	work := head
	result := make([]int, 0)
	for ; work != nil; work = work.Next {
		result = append([]int{work.Val}, result...)
	}
	return result
}
