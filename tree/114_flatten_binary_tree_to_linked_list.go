package tree

func flatten(root *TreeNode) {
	if nil == root {
		return
	}
	if nil == root.Left {
		flatten(root.Right)
	} else {
		rightChild := root.Right
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil
		if rightChild != nil {
			var cur *TreeNode
			for cur = root.Right; cur.Right != nil; cur = cur.Right {
			}
			cur.Right = rightChild
			flatten(rightChild)
		}
	}
}
