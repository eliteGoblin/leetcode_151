package sort

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pseudoHead := &ListNode{}
	pre := pseudoHead
	for l1 != nil && l2 != nil {
		var biggerNode **ListNode
		if l1.Val <= l2.Val {
			biggerNode = &l1
		} else {
			biggerNode = &l2
		}
		pre.Next = *biggerNode
		pre = pre.Next
		*biggerNode = (*biggerNode).Next
	}
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}
	return pseudoHead.Next
}
