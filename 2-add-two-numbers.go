package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	res := &ListNode{
		Val: 0, Next: &ListNode{},
	}
	cursor := res
	carry := 0

	// carry > 0 补足最高位
	for l1 != nil || l2 != nil || carry > 0 {
		cursor.Next = &ListNode{}
		cursor = cursor.Next

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cursor.Val = carry % 10
		carry /= 10
	}

	return res.Next
}