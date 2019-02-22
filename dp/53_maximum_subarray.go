package dp

// dp solution
func maxSubArray(nums []int) int {
	state := make([]int, len(nums))
	maxSum := nums[0]
	state[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		state[i] = max(state[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, state[i])
	}
	return maxSum
}

// naive solution

func maxSubArrayNaive(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxElement := nums[0]
	curSum := 0
	maxSum := 0
	for i := range nums {
		maxElement = max(maxElement, nums[i])
		if curSum+nums[i] > 0 {
			curSum += nums[i]
			maxSum = max(maxSum, curSum)
		} else {
			curSum = 0
		}
	}
	if maxElement < 0 {
		return maxElement
	}
	return max(maxElement, maxSum)
}

func max(arr ...int) int {
	res := arr[0]
	for i := range arr {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
