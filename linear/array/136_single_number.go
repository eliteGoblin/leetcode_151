package array

func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		res = res ^ nums[i]
	}
	return res
}
