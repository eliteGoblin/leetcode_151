package linklist

func partition(head *ListNode, x int) *ListNode {
	ltListHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	geListHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	preLTList := ltListHead
	preGEListHead := geListHead

	for cur := head; cur != nil; {
		if cur.Val < x {
			preLTList.Next = cur
			preLTList = cur
		} else {
			preGEListHead.Next = cur
			preGEListHead = cur
		}
		cur = cur.Next
	}
	preLTList.Next = geListHead.Next
	preGEListHead.Next = nil
	return ltListHead.Next
}
