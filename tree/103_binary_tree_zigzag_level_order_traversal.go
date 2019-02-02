package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	leftToRight := true
	zigzagTraverse(root, 0, leftToRight, &res)
	return res
}

func zigzagTraverse(root *TreeNode, level int, leftToRight bool, res *[][]int) {
	if nil == root {
		return
	}
	if level >= len(*res) {
		*res = append(*res, make([]int, 0))
	}

	if leftToRight {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		(*res)[level] = append([]int{root.Val}, (*res)[level][:]...)
	}
	leftToRight = !leftToRight
	zigzagTraverse(root.Left, level+1, leftToRight, res)
	zigzagTraverse(root.Right, level+1, leftToRight, res)
}
