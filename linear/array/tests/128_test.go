package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases128 = []struct {
	in     []int
	expect int
}{
	{
		in:     []int{100, 4, 200, 1, 3, 2},
		expect: 4,
	},
}

func TestLongestConsecutive(t *testing.T) {
	for i, tt := range testCases128 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.LongestConsecutive(tt.in)
			assert.True(t, res == tt.expect, "ecpect %d, got %d", tt.expect, res)
		})
	}
}
