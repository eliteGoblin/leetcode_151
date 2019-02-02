package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	} else if nil == p || nil == q {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
