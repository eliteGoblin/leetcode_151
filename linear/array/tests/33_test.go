package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	in     []int
	target int
	out    int
}{
	{
		in:     []int{4, 5, 6, 7, 0, 1, 2},
		target: 0,
		out:    4,
	},
	{
		in:     []int{4, 5, 6, 7, 0, 1, 2},
		target: 3,
		out:    -1,
	},
	{
		in:     []int{5, 1, 3},
		target: 3,
		out:    2,
	},
	{
		in:     []int{1, 3},
		target: 3,
		out:    1,
	},
	{
		in:     []int{3, 1},
		target: 1,
		out:    1,
	},
	{
		in:     []int{2, 5, 6, 0, 0, 1, 2},
		target: 0,
		out:    3,
	},
}

func TestSearch(t *testing.T) {
	for i, tt := range testCases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.Search(tt.in, tt.target)
			assert.True(t, res == tt.out, "ecpect %d, got %d", tt.out, res)
		})
	}
}
