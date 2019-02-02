package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"
)

var testCases15 = []struct {
	in     []int
	expect [][]int
}{
	{
		in: []int{-1, 0, 1, 2, -1, -4},
		expect: [][]int{
			[]int{-1, -1, 2},
			[]int{-1, 0, 1},
		},
	},
	{
		in: []int{3, 0, -2, -1, 1, 2},
	},
}

func TestThreeSum(t *testing.T) {
	for i, v := range testCases15 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.ThreeSum(v.in)
			// assert.Equal(t, v.expect, res, "ecpect %+v, got %+v", v.expect, res)
			fmt.Println(res)
		})
	}
}
