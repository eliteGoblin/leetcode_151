package tree

func generateTrees(n int) []*TreeNode {
	if 0 == n {
		return []*TreeNode{
			nil,
		}
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{
			nil,
		}
	}
	if start == end {
		return []*TreeNode{
			&TreeNode{
				Val: start,
			},
		}
	}
	var leftSubTree, rightSubTree []*TreeNode
	res := make([]*TreeNode, 0)
	for k := start; k <= end; k++ {
		leftSubTree = generate(start, k-1)
		rightSubTree = generate(k+1, end)
		for _, vLeft := range leftSubTree {
			for _, vRight := range rightSubTree {
				root := &TreeNode{
					Val:   k,
					Left:  vLeft,
					Right: vRight,
				}
				res = append(res, root)
			}
		}
	}
	return res
}
