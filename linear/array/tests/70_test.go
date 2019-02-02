package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases70 = []struct {
	in     int
	expect int
}{
	{
		in:     2,
		expect: 2,
	},
	{
		in:     5,
		expect: 8,
	},
}

func TestClimbStairs(t *testing.T) {
	for i, tt := range testCases70 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.ClimbStairs(tt.in)
			assert.Equal(t, tt.expect, res)
		})
	}
}
