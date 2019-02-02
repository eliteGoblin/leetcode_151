package tree

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left >= right {
		return left + 1
	}
	return right + 1
}