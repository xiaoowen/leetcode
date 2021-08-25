package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	tmp := head
	l := 0
	for tmp != nil {
		l++
		tmp = tmp.Next
	}

	pos := l / 2 + 1
	for i := 1; i<pos; i++ {
		head = head.Next
	}
	return head
}

// 快慢指针的方法
func middleNode1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}