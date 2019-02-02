package sort

// 1. 直接qsort
// 2. 双指针
func sortColors(nums []int) {
	red, blue := -1, len(nums)
	for i := 0; i < blue && red < blue; {
		switch nums[i] {
		case 0:
			red++
			// [0] 情况，red == i, red不应该再往前走, red不应超越i
			if red == i {
				i++
				break
			}
			nums[red], nums[i] = nums[i], nums[red]
		case 1:
			i++
		case 2:
			blue--
			nums[blue], nums[i] = nums[i], nums[blue]
		}
	}
}

// 3.两次partition: 先按0, 再按1
