package sort

// naive approach
// func mergeKLists(lists []*ListNode) *ListNode {
// 	pseudoHead := &ListNode{}
// 	preNode := pseudoHead
// 	for node, end := nextNode(lists); !end; {
// 		preNode.Next = node
// 		preNode = preNode.Next
// 	}
// 	return pseudoHead.Next
// }

// func nextNode(lists []*ListNode) (nextNode *ListNode, isEnd bool) {
// 	isEnd = true
// 	var maxNode *ListNode
// 	var maxIndex int
// 	for i, v := range lists {
// 		if nil != v {
// 			isEnd = false
// 			if maxNode == nil || v.Val > maxNode.Val {
// 				maxNode = v
// 				maxIndex = i
// 			}
// 		}
// 	}
// 	if !isEnd {
// 		lists[maxIndex] = lists[maxIndex].Next
// 	}
// 	return maxNode, isEnd
// }
