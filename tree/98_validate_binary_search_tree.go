package tree

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func isValidBST(root *TreeNode) bool {
	return isValidBSTRange(root, MinInt, MaxInt)
}

func isValidBSTRange(root *TreeNode, lowerBound int, upperBound int) bool {
	if nil == root {
		return true
	}
	return (root.Val > lowerBound) &&
		(root.Val < upperBound) &&
		isValidBSTRange(root.Left, lowerBound, root.Val) &&
		isValidBSTRange(root.Right, root.Val, upperBound)
}
