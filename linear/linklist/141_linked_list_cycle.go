package linklist

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
	}
	return false
}
