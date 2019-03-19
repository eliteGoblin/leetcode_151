package review_151

// 自己recursive方法
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}
	start, last := 0, len(nums)-1
	mid := last / 2
	if nums[mid] == target {
		return true
	}
	for start != mid && nums[start] == nums[mid] {
		start++
	}
	if start == mid {
		return search(nums[mid+1:], target)
	}
	for last != mid && nums[last] == nums[mid] {
		last--
	}
	if last == mid {
		return search(nums[:mid], target)
	}
	if nums[mid] > nums[start] {
		if target < nums[mid] && target >= nums[start] {
			return search(nums[:mid], target)
		} else {
			return search(nums[mid+1:], target)
		}
	} else {
		if target > nums[mid] && target <= nums[last] {
			return search(nums[mid+1:], target)
		} else {
			return search(nums[:mid], target)
		}
	}
}
