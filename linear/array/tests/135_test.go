package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases135 = []struct {
	in     []int
	expect int
}{
	{
		in:     []int{1, 0, 2},
		expect: 5,
	},
	{
		in:     []int{1, 2, 2},
		expect: 4,
	},
	{
		in:     []int{1, 3, 2, 2, 1},
		expect: 7,
	},
}

func TestCandy(t *testing.T) {
	for i, tt := range testCases135 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.Candy(tt.in)
			assert.Equal(t, tt.expect, res)
		})
	}
}
