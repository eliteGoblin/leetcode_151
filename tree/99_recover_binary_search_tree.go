package tree

var path []*TreeNode

func recoverTree(root *TreeNode) {
	path = make([]*TreeNode, 0)
	inorderTraversalStore(root)
	var i, j int
	for i = 0; i < len(path)-1; i++ {
		if path[i].Val > path[i+1].Val {
			break
		}
	}
	for j = len(path) - 1; j > 0; j-- {
		if path[j].Val < path[j-1].Val {
			break
		}
	}
	path[i].Val, path[j].Val = path[j].Val, path[i].Val
}

func inorderTraversalStore(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalStore(root.Left)
	path = append(path, root)
	inorderTraversalStore(root.Right)
}
