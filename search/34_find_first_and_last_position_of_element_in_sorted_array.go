package search

// type predicate func(i, target int) bool

// func searchRange(nums []int, target int) []int {
// 	isStart := func(i int, target int) bool {
// 		if nums[i] != target {
// 			return false
// 		}
// 		if i == 0 {
// 			return true
// 		}
// 		if nums[i-1] != target {
// 			return true
// 		}
// 		return false
// 	}
// 	isEnd := func(i int, target int) bool {
// 		if nums[i] != target {
// 			return false
// 		}
// 		if i == len(nums)-1 {
// 			return true
// 		}
// 		if nums[i+1] != target {
// 			return true
// 		}
// 		return false
// 	}
// 	binarySearch := func(target int, findStart bool) int {
// 		start, end := 0, len(nums)-1
// 		mid := start + (end-start)/2
// 		if nums[mid] == target {
// 			if findStart {
// 				if isStart()
// 				end = mid - 1
// 			} else {

// 			}
// 		}
// 	}
// 	start := binarySearch(target, isStart)
// 	if start == -1 {
// 		return []int{-1, -1}
// 	}
// 	end := binarySearch(target, isEnd)
// 	return []int{start, end}
// }

func searchRange(nums []int, target int) []int {
	fst, last := -1, -1
	// find first
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			if 0 == mid || nums[mid-1] < target {
				fst = mid
				break
			} else {
				end = mid - 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if fst == -1 {
		return []int{-1, -1}
	}
	start, end = 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				last = mid
				break
			} else {
				start = mid + 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return []int{fst, last}
}
