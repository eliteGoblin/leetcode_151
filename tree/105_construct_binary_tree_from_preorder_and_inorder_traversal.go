package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	leftIndexEnd := 0
	for ; leftIndexEnd < len(inorder); leftIndexEnd++ {
		if inorder[leftIndexEnd] == root.Val {
			break
		}
	}
	root.Left = buildTree(preorder[1:leftIndexEnd+1], inorder[:leftIndexEnd])
	root.Right = buildTree(preorder[leftIndexEnd+1:], inorder[leftIndexEnd+1:])
	return root
}
