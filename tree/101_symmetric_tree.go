package tree

//  best solution
func isSymmetricBest(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return areTwoNodeSymmetric(root.Left, root.Right)
}

func areTwoNodeSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		areTwoNodeSymmetric(left.Left, right.Right) &&
		areTwoNodeSymmetric(left.Right, right.Left)
}

// my invert + isSame solution
func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	flipRoot := genFlipTree(root)
	return isSameTree(root, flipRoot)
}

func genFlipTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := &TreeNode{
		Val: root.Val,
	}
	node.Left = genFlipTree(root.Right)
	node.Right = genFlipTree(root.Left)
	return node
}
