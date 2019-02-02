package tests

import (
	"testing"
)

var testCases4 = []struct {
	nums1  []int
	nums2  []int
	expect float64
}{
	{
		nums1:  []int{1, 3},
		nums2:  []int{2},
		expect: 2.0,
	},
	// {
	// 	nums1:  []int{1, 2},
	// 	nums2:  []int{3, 4},
	// 	expect: 2.5,
	// },
}

func TestFindMedianSortedArrays(t *testing.T) {
	// for i, tt := range testCases4 {
	// 	t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
	// 		res := array.FindMedianSortedArrays(tt.nums1, tt.nums2)
	// 		assert.True(t, math.Abs(res-tt.expect) <= 1e-6, "ecpect %f, got %f", tt.expect, res)
	// 	})
	// }
}
