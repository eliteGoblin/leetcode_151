package tree

func numTrees(n int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return countTree(arr)
}

func countTree(arr []int) int {
	res := 0
	if len(arr) == 1 {
		return 1
	}
	for i, v := range arr {
		res += countTree(arr[:i])
		if i+1 < len(arr) {
			res += countTree(arr[i+1:])
		}
	}
	return res
}
