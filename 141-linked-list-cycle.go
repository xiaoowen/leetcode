package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	scaned := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := scaned[head]; ok {
			return true
		}
		scaned[head] = true
		head = head.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
