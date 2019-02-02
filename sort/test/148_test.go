package test

import (
	"fmt"
	"leetcode/sort"
	"testing"
)

func TestSortList(t *testing.T) {
	lst := genList([]int{4, 2, 1, 3})
	lst = sort.SortList(lst)
}

func genList(arr []int) *sort.ListNode {
	pseudoHead := &sort.ListNode{}
	pre := pseudoHead
	for _, v := range arr {
		pre.Next = &sort.ListNode{
			Val: v,
		}
		pre = pre.Next
	}
	return pseudoHead.Next
}

func printList(head *sort.ListNode) {
	for head != nil {
		fmt.Printf(" %d", head.Val)
		head = head.Next
	}
	fmt.Println("")
}
