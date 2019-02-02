package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases48 = []struct {
	in     [][]int
	expect [][]int
}{
	{
		in: [][]int{
			{
				1, 2, 3,
			},
			{
				4, 5, 6,
			},
			{
				7, 8, 9,
			},
		},
		expect: [][]int{
			{
				7, 4, 1,
			},
			{
				8, 5, 2,
			},
			{
				9, 6, 3,
			},
		},
	},
}

func TestRotate(t *testing.T) {
	for i, tt := range testCases48 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			array.Rotate(tt.in)
			assert.Equal(t, tt.in, tt.expect, "ecpect %d, got %d", tt.expect, tt.in)
		})
	}
}
