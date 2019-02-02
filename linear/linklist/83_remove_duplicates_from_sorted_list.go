package linklist

func deleteDuplicates(head *ListNode) *ListNode {
	notDupNode := head
	for notDupNode != nil {
		firstNotDupNode := notDupNode.Next
		for firstNotDupNode != nil && firstNotDupNode.Val == notDupNode.Val {
			firstNotDupNode = firstNotDupNode.Next
		}
		notDupNode.Next = firstNotDupNode
		notDupNode = firstNotDupNode
	}
	return head
}
