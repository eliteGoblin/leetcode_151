package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases27 = []struct {
	nums   []int
	val    int
	expect int
}{
	{
		nums:   []int{3, 2, 2, 3},
		val:    3,
		expect: 2,
	},
	{
		nums:   []int{1},
		val:    1,
		expect: 0,
	},
}

func TestRemoveElement(t *testing.T) {
	for i, tt := range testCases27 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.RemoveElement(tt.nums, tt.val)
			assert.True(t, res == tt.expect, "ecpect %d, got %d", tt.expect, res)
		})
	}
}
