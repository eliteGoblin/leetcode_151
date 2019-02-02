package template

func QSort(arr []int) {
	if len(arr) > 1 {
		i := partition(arr)
		QSort(arr[:i])
		QSort(arr[i+1:])
	}
}

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	index := -1
	for i := range arr {
		if arr[i] <= pivot {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	return index
}

// reference:
// [QuickSort](https://www.geeksforgeeks.org/quick-sort/)
