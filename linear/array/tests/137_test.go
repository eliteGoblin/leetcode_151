package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases137 = []struct {
	in     []int
	expect int
}{
	// {
	// 	in:     []int{2, 2, 3, 2},
	// 	expect: 3,
	// },
	// {
	// 	in:     []int{0, 1, 0, 1, 0, 1, 99},
	// 	expect: 99,
	// },
	{
		in:     []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2},
		expect: -4,
	},
}

func TestSingleNumberII(t *testing.T) {
	for i, tt := range testCases137 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.SingleNumberII(tt.in)
			assert.Equal(t, tt.expect, res)
		})
	}
}
