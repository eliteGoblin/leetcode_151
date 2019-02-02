package array

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	currentDedupIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[currentDedupIndex] {
			currentDedupIndex++
			nums[currentDedupIndex] = nums[i]
		}
	}
	return currentDedupIndex + 1
}
