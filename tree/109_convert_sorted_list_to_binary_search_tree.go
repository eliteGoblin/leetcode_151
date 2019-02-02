package tree

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if nil == head {
		return nil
	}
	return sortedListRangeToBST(head, nil)
}

func sortedListRangeToBST(start, end *TreeNode) *TreeNode {
	if start == end {
		return nil
	}
	midNode := getListMid(start, end)
	root := &TreeNode{
		Val: midNode.Val,
	}
	root.Left = sortedListRangeToBST(start, mid)
	root.Right = sortedListRangeToBST(mid.Next, end)
	return root
}
func getListMid(head, end *ListNode) *ListNode {
	if end == head {
		return nil
	}
	fast, slow := head, head
	for fast.Next != end && fast.Next.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
