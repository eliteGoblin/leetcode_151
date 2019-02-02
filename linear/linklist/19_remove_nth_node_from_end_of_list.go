package linklist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	pseudoHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := pseudoHead
	var tail *ListNode
	steps := n + 1
	for tail = pseudoHead; tail != nil && steps > 0; {
		tail = tail.Next
		steps--
	}
	if steps == 0 {
		if nil == tail {
			return pseudoHead.Next.Next
		}
	} else {
		return head
	}
	for tail != nil {
		tail = tail.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return pseudoHead.Next
}
