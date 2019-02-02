package sort

func SortList(head *ListNode) *ListNode {
	len := getListLen(head)
	if len <= 1 {
		return head
	}
	return mergeSortList(head, len)
}

func getListLen(head *ListNode) int {
	res := 0
	for head != nil {
		res++
		head = head.Next
	}
	return res
}

func mergeSortList(head *ListNode, len int) *ListNode {
	if len == 1 {
		head.Next = nil
		return head
	}

	secListNode := advance(head, len/2)
	left := mergeSortList(head, len/2)
	right := mergeSortList(secListNode, len-len/2)
	return merge2SortedList(left, right)
}

func advance(head *ListNode, step int) *ListNode {
	for step > 0 {
		head = head.Next
		step--
	}
	return head
}

func merge2SortedList(la, lb *ListNode) *ListNode {
	pseudoHead := &ListNode{}
	pre := pseudoHead
	for la != nil && lb != nil {
		var smaller *ListNode
		if la.Val <= lb.Val {
			smaller = la
			la = la.Next
		} else {
			smaller = lb
			lb = lb.Next
		}
		pre.Next = smaller
		pre = pre.Next
	}
	if la != nil {
		pre.Next = la
	}
	if lb != nil {
		pre.Next = lb
	}
	return pseudoHead.Next
}
