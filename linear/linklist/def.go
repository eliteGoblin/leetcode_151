package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Slice2List(arr []int) *ListNode {
	pseudoHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	pre := pseudoHead
	for _, v := range arr {
		pre.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		pre = pre.Next
	}
	pre.Next = nil
	return pseudoHead.Next
}

func List2Slice(head *ListNode) []int {
	res := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}
	return res
}

func PrintList(head *ListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("% 02d", cur.Val)
	}
	fmt.Println()
}
