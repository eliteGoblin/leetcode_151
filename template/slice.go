package template

// tmp slice and copy back
// func Insert(arr []int, pos int, value int) []int {
// 	return append(arr[:pos], append([]int{value}, arr[pos:]...)...)
// }

func Insert(arr []int, pos int, value int) []int {
	s := append(arr, 0)
	copy(s[pos+1:], s[pos:])
	s[pos] = value
	return s
}
