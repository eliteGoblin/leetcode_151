package sort

// {4,2,1,3}
func insertionSortList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	pseudoHead := &ListNode{}
	var pre *ListNode
	for cur := head; cur != nil; {
		pre = findInsertPre(pseudoHead, cur.Val)
		sortedNext := pre.Next
		next := cur.Next
		pre.Next = cur
		cur.Next = sortedNext
		cur = next
	}
	return pseudoHead.Next
}

func findInsertPre(head *ListNode, newValue int) (preNode *ListNode) {
	for preNode = head; preNode.Next != nil; preNode = preNode.Next {
		if preNode.Next.Val >= newValue {
			break
		}
	}
	return preNode
}
