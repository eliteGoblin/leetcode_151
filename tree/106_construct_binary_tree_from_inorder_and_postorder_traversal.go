package tree

func buildTreeII(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	leftTreeLimit := 0
	for ; leftTreeLimit < len(inorder); leftTreeLimit++ {
		if inorder[leftTreeLimit] == root.Val {
			break
		}
	}
	root.Left = buildTreeII(inorder[:leftTreeLimit], postorder[:leftTreeLimit])
	root.Right = buildTreeII(inorder[leftTreeLimit+1:], postorder[leftTreeLimit:len(postorder)-1])
	return root
}
