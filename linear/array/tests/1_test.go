package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases1 = []struct {
	in     []int
	target int
	expect []int
}{
	{
		in:     []int{2, 7, 11, 15},
		target: 9,
		expect: []int{0, 1},
	},
	{
		in:     []int{3, 2, 4},
		target: 6,
		expect: []int{1, 2},
	},
}

func TestTwoSum(t *testing.T) {
	for i, tt := range testCases1 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.TwoSum(tt.in, tt.target)
			assert.Equal(t, tt.expect, res)
		})
	}
}
