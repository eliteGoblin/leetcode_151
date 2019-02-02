package linklist

func ReverseKGroupRecursive(head *ListNode, k int) *ListNode {
	end, ok := forwardKStep(head, k)
	if !ok {
		return head
	}
	pseudoHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	for cur := head; cur != end; {
		next := cur.Next
		cur.Next = pseudoHead.Next
		pseudoHead.Next = cur
		cur = next
	}
	head.Next = ReverseKGroupRecursive(end, k)
	return pseudoHead.Next
}
