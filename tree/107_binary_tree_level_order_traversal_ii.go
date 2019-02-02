package tree

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	traverse(root, 0, res)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func traverse(root *TreeNode, level int, res [][]int) {
	if nil == root {
		return
	}
	if level >= len(res) {
		res = append(res, make([]int, 0))
	}
	res[level] = append(res[level], root.Val)
	traverse(root.Left, level+1, res)
	traverse(root.Right, level+1, res)
}
