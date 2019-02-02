package tree

func sumNumbers(root *TreeNode) int {
	res := 0
	path := 0
	rootToPath(root, &res, path)
	return res
}

func rootToPath(root *TreeNode, res *int, pathSum int) {
	if nil == root {
		return
	}
	pathSum = pathSum*10 + root.Val
	if nil == root.Left && nil == root.Right {
		*res += pathSum
		return
	}
	rootToPath(root.Left, res, pathSum)
	rootToPath(root.Right, res, pathSum)
}
