package array

func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		mp[v] = i
	}
	for i, v := range nums {
		another := target - v
		if j, ok := mp[another]; ok {
			if j != i {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
