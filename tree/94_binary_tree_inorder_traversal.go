package tree

import (
	"leetcode/stack"
)

func inorderTraversal(root *TreeNode) []int {
	stack := stack.NewStack()
	cur := root
	res := make([]int, 0)
	for cur != nil || !stack.Empty() {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			e, _ := stack.Pop()
			nextNode := e.(*TreeNode)
			res = append(res, nextNode.Val)
			cur = nextNode.Right
		}
	}
	return res
}
