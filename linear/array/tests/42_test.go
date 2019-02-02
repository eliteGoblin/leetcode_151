package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases42 = []struct {
	in     []int
	expect int
}{
	{
		in:     []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		expect: 2,
	},
}

func TestTrap(t *testing.T) {
	for i, tt := range testCases42 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.Trap(tt.in)
			assert.Equal(t, tt.expect, res)
		})
	}
}
