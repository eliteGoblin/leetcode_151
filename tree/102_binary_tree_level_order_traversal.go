package tree

func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return [][]int{}
	}
	queue := NewQueue()
	queue.PushBack(root)
	count := queue.Len()
	res := make([][]int, 0)
	layer := make([]int, 0)
	for {
		if count > 0 {
			v, _ := queue.PopFront()
			e := v.(*TreeNode)
			layer = append(layer, e.Val)
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
			count--
		} else {
			res = append(res, layer)
			if queue.Empty() {
				break
			}
			count = queue.Len()
			layer = make([]int, 0)
		}
	}
	return res
}
