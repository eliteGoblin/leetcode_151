package tree

import (
	"leetcode/utils"
)

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	_, res := isBalancedAndGetHeight(root)
	return res
}

func isBalancedAndGetHeight(root *TreeNode) (height int, isBalanced bool) {
	if nil == root {
		return 0, true
	}
	heightL, okL := isBalancedAndGetHeight(root.Left)
	heightR, okR := isBalancedAndGetHeight(root.Right)
	if okL && okR {
		return utils.Max(heightL, heightR) + 1, abs(heightL, heightR) <= 1
	}
	return -1, false
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
