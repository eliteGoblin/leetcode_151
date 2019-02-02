package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases31 = []struct {
	nums   []int
	expect []int
}{
	{
		nums:   []int{1, 2, 3},
		expect: []int{1, 3, 2},
	},
	{
		nums:   []int{3, 2, 1},
		expect: []int{1, 2, 3},
	},
	{
		nums:   []int{1, 1, 5},
		expect: []int{1, 5, 1},
	},
	{
		nums:   []int{6, 8, 7, 4, 3, 2},
		expect: []int{7, 2, 3, 4, 6, 8},
	},
}

func TestNextPermutation(t *testing.T) {
	for i, tt := range testCases31 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			array.NextPermutation(tt.nums)
			assert.Equal(t, tt.nums, tt.expect, "ecpect %d, got %d", tt.expect, tt.nums)
		})
	}
}
