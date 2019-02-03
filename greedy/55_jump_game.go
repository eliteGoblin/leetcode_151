package greedy

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	lastIndex := len(nums) - 1
	limit := nums[0]
	for i := 1; i <= limit; {
		if i >= lastIndex {
			return true
		}
		if i+nums[i] > limit {
			limit = i + nums[i]
		}
	}
	return false
}

// 超时
func canJumpDFS(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	for i := 1; i <= nums[0]; i++ {
		if i >= len(nums) {
			break
		}
		if res := canJumpDFS(nums[i:]); res {
			return true
		}
	}
	return false
}
