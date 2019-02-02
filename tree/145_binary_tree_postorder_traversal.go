package tree

import (
	"leetcode/stack"
)

func postorderTraversal(root *TreeNode) []int {
	preOrderStack := stack.NewStack()
	resultStack := stack.NewStack()
	cur := root
	res := make([]int, 0)
	for cur != nil || !preOrderStack.Empty() {
		if cur != nil {
			resultStack.Push(cur)
			preOrderStack.Push(cur.Left)
			cur = cur.Right
		} else {
			e, _ := preOrderStack.Pop()
			cur = e.(*TreeNode)
		}
	}
	for !resultStack.Empty() {
		e, _ := resultStack.Pop()
		res = append(res, e.(*TreeNode).Val)
	}
	return res
}
