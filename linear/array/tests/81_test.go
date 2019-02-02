package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases81 = []struct {
	in     []int
	target int
	expect bool
}{
	{
		in:     []int{2, 5, 6, 0, 0, 1, 2},
		target: 0,
		expect: true,
	},
	{
		in:     []int{2, 5, 6, 0, 0, 1, 2},
		target: 3,
		expect: false,
	},
	{
		in:     []int{1, 1, 3, 1},
		target: 3,
		expect: true,
	},
}

func TestSearchII(t *testing.T) {
	for i, tt := range testCases81 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.SearchII(tt.in, tt.target)
			assert.True(t, res == tt.expect, "ecpect %t, got %t", tt.expect, res)
		})
	}
}
