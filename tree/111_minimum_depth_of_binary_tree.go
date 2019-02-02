package tree

const INT_MAX = int((^uint(0)) >> 1)

func minDepth(root *TreeNodel) int {
	return minDepthRecur(root, true)
}

func minDepthRecur(root *TreeNode, isRoot bool) int {
	if nil == root {
		if isRoot {
			return 0
		}
		return INT_MAX
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := minDepthRecur(root.Left, false)
	right := minDepthRecur(root.Right, false)
	if left < right {
		return left + 1
	}
	return right + 1
}
