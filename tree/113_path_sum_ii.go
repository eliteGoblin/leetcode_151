package tree

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	pathSumRecur(root, sum, []int{}, &res)
	return res
}

func pathSumRecur(root *TreeNode, sum int, curPath []int, res *[][]int) {
	if nil == root {
		return
	}
	curPath = append(curPath, root.Val)
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			(*res) = append(*res, curPath)
		}
		return
	}
	pathSumRecur(root.Left, sum-root.Val, curPath, res)
	pathSumRecur(root.Right, sum-root.Val, curPath, res)
	return
}
