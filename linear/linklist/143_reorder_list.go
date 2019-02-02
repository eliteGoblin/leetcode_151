package linklist

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	firstHalf, secondHalf := splitList(head)
	secondHalf = reverseList(secondHalf)
	mergeList(firstHalf, secondHalf)
}

// fast, slow pointer to get middle
func splitList(head *ListNode) (firstHalf, secondHalf *ListNode) {
	len := listLen(head)
	firstHalf = head
	end := head
	if len%2 == 0 {
		end = moveList(head, len/2-1)
	} else {
		end = moveList(head, len/2)
	}
	secondHalf = end.Next
	end.Next = nil
	return
}

func listLen(head *ListNode) int {
	res := 0
	for cur := head; cur != nil; cur = cur.Next {
		res++
	}
	return res
}

func moveList(head *ListNode, step int) *ListNode {
	var cur *ListNode
	for cur = head; cur != nil && step > 0; {
		cur = cur.Next
		step--
	}
	return cur
}

func reverseList(head *ListNode) *ListNode {
	pseudoHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	pre := pseudoHead
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre.Next
		pre.Next = cur

		cur = next
	}
	return pseudoHead.Next
}

func mergeList(firstHalf, secondHalf *ListNode) {
	pseudoHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	pre := pseudoHead
	for firstHalf != nil && secondHalf != nil {
		nextFirsrtHalf := firstHalf.Next
		pre.Next = firstHalf
		firstHalf.Next = secondHalf
		pre = secondHalf
		firstHalf = nextFirsrtHalf
		secondHalf = secondHalf.Next
	}
	if firstHalf != nil {
		pre.Next = firstHalf
		pre = pre.Next
	}
	pre.Next = nil
}
