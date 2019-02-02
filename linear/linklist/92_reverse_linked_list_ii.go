package linklist

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n <= m {
		return head
	}
	pseudoHead := &ListNode{
		Val:  0,
		Next: head,
	}
	preNode := pseudoHead
	for i := 1; preNode != nil && i < m; i++ {
		preNode = preNode.Next
	}
	newLastNode := preNode.Next
	cur := preNode.Next
	for i := m; i <= n; i++ {
		nextNode := cur.Next
		cur.Next = preNode.Next
		preNode.Next = cur
		cur = nextNode
	}
	newLastNode.Next = cur
	return pseudoHead.Next
}
