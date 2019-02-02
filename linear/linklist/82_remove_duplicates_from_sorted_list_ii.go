package linklist

func deleteDuplicatesII(head *ListNode) *ListNode {
	pseudoHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	preNode := pseudoHead
	cur := head
	for cur != nil {
		if node, dup := isDupAndGetNextDifferentNode(cur); dup {
			cur = node
		} else {
			preNode.Next = cur
			preNode = cur
			cur = cur.Next
		}
	}
	preNode.Next = nil
	return pseudoHead.Next
}

func isDupAndGetNextDifferentNode(cur *ListNode) (next *ListNode, isDup bool) {
	next = cur.Next
	isDup = false
	for next != nil && next.Val == cur.Val {
		isDup = true
		next = next.Next
	}
	return next, isDup
}
