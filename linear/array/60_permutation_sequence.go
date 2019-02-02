package array

import (
	"fmt"
)

func GetPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	for i := 1; i < k; i++ {
		NextPermutation(nums)
	}
	res := ""
	for i := 0; i < n; i++ {
		res += fmt.Sprintf("%d", nums[i])
	}
	return res
}
