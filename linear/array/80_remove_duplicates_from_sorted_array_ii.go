package array

func RemoveDuplicatesII(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	index := 0
	curRepTimes := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
			curRepTimes = 0
		} else if 0 == curRepTimes {
			index++
			nums[index] = nums[i]
			curRepTimes++
		} else {
			curRepTimes++
		}
	}
	return index + 1
}
