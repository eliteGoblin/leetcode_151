package linklist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	pseudoHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	pre := pseudoHead
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		l1.Val = sum
		pre.Next = l1
		pre = pre.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	var linkLeft *ListNode
	if l1 != nil {
		linkLeft = l1
	} else {
		linkLeft = l2
	}
	for linkLeft != nil {
		sum := linkLeft.Val + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		linkLeft.Val = sum
		pre.Next = linkLeft
		pre = pre.Next
		linkLeft = linkLeft.Next
	}
	if carry > 0 {
		pre.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	} else {
		pre.Next = nil
	}
	return pseudoHead.Next
}
