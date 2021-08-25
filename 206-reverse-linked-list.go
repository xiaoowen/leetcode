package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		// 将prev给cur的next
		cur.Next = prev
		// cur跳到下一个节点，prev
		prev = cur
		cur = next
	}
	// 返回原链表最后一个节点
	return prev
}
