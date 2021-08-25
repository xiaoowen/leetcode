package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	// slow and fast pointer
	vh := &ListNode{Next: head}

	slow, fast := vh, vh

	// fast先走n+1，然后与slow一起走
	for i := 0; i < n+1; i ++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return vh.Next
}
