package tests

import (
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases16 = []struct {
	in     []int
	target int
	expect int
}{
	// {
	// 	in:     []int{-1, 2, 1, -4},
	// 	target: 1,
	// 	expect: 2,
	// },
	{
		in:     []int{0, 2, 1, -3},
		target: 1,
		expect: 0,
	},
}

func Test3SumClosest(t *testing.T) {
	for _, v := range testCases16 {
		t.Run("test", func(t *testing.T) {
			res := array.ThreeSumClosest(v.in, v.target)
			assert.True(t, v.expect == res, "expect %d, got %d", v.expect, res)
		})
	}
}
