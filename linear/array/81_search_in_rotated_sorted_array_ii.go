package array

func SearchII(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if target == nums[mid] {
			return true
		}
		for mid < end && nums[mid] == nums[end] {
			end--
		}
		if nums[mid] <= nums[end] {
			// 右边有序
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			// 左边有序
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return false
}
