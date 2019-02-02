package array

import (
	"fmt"
	"unsafe"
)

func SingleNumberII(nums []int) int {
	if len(nums) < 4 {
		return nums[0]
	}
	bits := int(unsafe.Sizeof(nums[0])) * 8
	fmt.Println(bits)
	var res int
	for i := 0; i < bits; i++ {
		sum := 0
		for _, v := range nums {
			sum += (v >> uint(i)) & 1
		}
		res |= (sum % 3) << uint(i)
	}
	return res
}
