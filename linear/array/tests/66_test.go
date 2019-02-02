package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases66 = []struct {
	in     []int
	expect []int
}{
	{
		in:     []int{1, 2, 3},
		expect: []int{1, 2, 4},
	},
	{
		in:     []int{9, 9, 9},
		expect: []int{1, 0, 0, 0},
	},
}

func TestPlusOne(t *testing.T) {
	for i, tt := range testCases66 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.PlusOne(tt.in)
			assert.Equal(t, tt.expect, res)
		})
	}
}
