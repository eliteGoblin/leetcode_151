package linklist

func ReverseKGroupInteration(head *ListNode, k int) *ListNode {
	pseudoHead := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := pseudoHead
	begin := head
	end, ok := forwardKStep(begin, k)
	for ok {
		pre = reverseKNodes(pre, begin, end)
		begin = pre.Next
		end, ok = forwardKStep(begin, k)
	}
	return pseudoHead.Next
}

func forwardKStep(pre *ListNode, k int) (*ListNode, bool) {
	for pre != nil && k > 0 {
		pre = pre.Next
		k--
	}
	if k == 0 {
		return pre, true
	}
	return &ListNode{}, false
}

func reverseKNodes(pre, begin, end *ListNode) (nextPre *ListNode) {
	pseudoHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := begin
	for cur != end {
		next := cur.Next
		cur.Next = pseudoHead.Next
		pseudoHead.Next = cur
		cur = next
	}
	pre.Next = pseudoHead.Next
	begin.Next = end
	return begin
}

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	if head == nil || k <= 1 {
// 		return head
// 	}
// 	pseudoHead := &ListNode{
// 		Val:  -1,
// 		Next: head,
// 	}
// 	pre := pseudoHead
// 	var ok = true
// 	for ok {
// 		pre, ok = reverseNextKNode(pre, k)
// 	}
// 	return pseudoHead.Next
// }

// func reverseNextKNode(pre *ListNode, k int) (nextPre *ListNode, ok bool) {
// 	if nil == pre && pre.Next == nil {
// 		return &ListNode{}, false
// 	}
// 	pseudoHead := &ListNode{
// 		Val:  -1,
// 		Next: nil,
// 	}
// 	cur := pre.Next
// 	newTail := pre.Next
// 	for cur != nil && k > 0 {
// 		tmp := cur.Next
// 		cur.Next = pseudoHead.Next
// 		pseudoHead.Next = cur

// 		cur = tmp
// 		k--
// 	}
// 	if cur == nil {
// 		if k > 0 {
// 			return &ListNode{}, false
// 		}
// 	}
// 	newTail.Next = cur
// 	pre.Next = pseudoHead.Next
// 	return newTail, true
// }
