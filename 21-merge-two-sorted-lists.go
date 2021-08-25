package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 递归
	var newl *ListNode
	if l1.Val <= l2.Val {
		newl = l1
		newl.Next = mergeTwoLists(l1.Next, l2)
	} else {
		newl = l2
		newl.Next = mergeTwoLists(l1, l2.Next)
	}
	return newl
}