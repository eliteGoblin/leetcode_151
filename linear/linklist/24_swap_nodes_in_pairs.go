package linklist

func swapPairs(head *ListNode) *ListNode {
	pseudoHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := pseudoHead
	for a, b, ok := getNext2Nodes(pre); ok; a, b, ok = getNext2Nodes(pre) {
		pre.Next = b
		a.Next = b.Next
		b.Next = a
		pre = a
	}
	return pseudoHead.Next
}

func getNext2Nodes(head *ListNode) (a *ListNode, b *ListNode, ok bool) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil, nil, false
	}
	return head.Next, head.Next.Next, true
}
