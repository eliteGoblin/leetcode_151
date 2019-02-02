package tree

import (
	"leetcode/stack"
)

func preorderTraversal(root *TreeNode) []int {
	stack := stack.NewStack()
	if nil == root {
		return []int{}
	}
	res := make([]int, 0)
	cur := root
	for cur != nil || !stack.Empty() {
		if cur != nil {
			res = append(res, cur.Val)
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			cur = cur.Left
		} else {
			e, _ := stack.Pop()
			cur = e.(*TreeNode)
		}
	}
	return res
}
