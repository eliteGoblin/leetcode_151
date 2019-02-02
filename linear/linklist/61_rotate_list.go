package linklist

import "fmt"

// 1 2 3 4 5 k = 2
func rotateRight(head *ListNode, k int) *ListNode {
	if k <= 0 || head == nil {
		return head
	}
	var len int
	for cur := head; cur != nil; cur = cur.Next {
		len++
	}
	if k >= len {
		k = k % len
	}
	pseudoHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := pseudoHead
	for i := 0; i < len-k; i++ {
		pre = pre.Next
	}
	newHead := pre.Next
	fmt.Println(newHead.Val)
	pre.Next = nil
	pseudoHead.Next = newHead
	for pre = newHead; pre.Next != nil; pre = pre.Next {
	}
	pre.Next = head
	return pseudoHead.Next
}
