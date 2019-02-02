package sort

import (
	"container/heap"
)

type PriorityQueue []*ListNode

func (selfPtr *PriorityQueue) Less(i, j int) bool {
	return (*selfPtr)[i].Val < (*selfPtr)[j].Val
}

func (selfPtr *PriorityQueue) Len() int {
	return len(*selfPtr)
}

func (selfPtr *PriorityQueue) Swap(i, j int) {
	(*selfPtr)[i], (*selfPtr)[j] = (*selfPtr)[j], (*selfPtr)[i]
}

func (selfPtr *PriorityQueue) Push(x interface{}) {
	*selfPtr = append(*selfPtr, x.(*ListNode))
}

func (selfPtr *PriorityQueue) Pop() interface{} {
	if len(*selfPtr) > 0 {
		item := (*selfPtr)[len(*selfPtr)-1]
		*selfPtr = (*selfPtr)[0 : len(*selfPtr)-1]
		return item
	}
	return &ListNode{}
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	for _, head := range lists {
		for head != nil {
			pq = append(pq, head)
			head = head.Next
		}
	}
	heap.Init(&pq)
	len := len(pq)
	pseudoHead := &ListNode{}
	pre := pseudoHead
	for i := 0; i < len; i++ {
		node := heap.Pop(&pq).(*ListNode)
		pre.Next = node
		pre = pre.Next
	}
	pre.Next = nil
	return pseudoHead.Next
}
