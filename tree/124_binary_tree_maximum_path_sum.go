package tree

import (
	"leetcode/utils"
)

func maxPathSum(root *TreeNode) int {
	res := 0
	maxPathRecur(root, &res)
	return res
}

func maxPathRecur(root *TreeNode, res *int) int {
	if nil == root {
		return 0
	}
	maxLeft := utils.Max(maxPathRecur(root.Left, res), 0)
	maxRight := utils.Max(maxPathRecur(root.Right, res), 0)
	res = utils.Max(res, maxLeft+root.Val+maxRight)
	return utils.Max(maxLeft+root.Val, maxRight+root.Val)
}
