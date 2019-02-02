package tree


func hasPathSum(root *TreeNode, sum int) bool {
	if nil == root {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) ||
	hasPathSum(root.Right, sum-root.Val) ||
}
