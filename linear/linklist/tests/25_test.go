package tests

import (
	"fmt"
	"leetcode/linear/linklist"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases25 = []struct {
	in     []int
	k      int
	expect []int
}{
	{
		in:     []int{1, 2, 3, 4, 5},
		k:      2,
		expect: []int{2, 1, 4, 3, 5},
	},
}

func TestReverseKGroupm(t *testing.T) {
	for i, v := range testCases25 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := linklist.ReverseKGroupRecursive(linklist.Slice2List(v.in), v.k)
			outSlice := linklist.List2Slice(res)
			assert.Equal(t, v.expect, outSlice, "ecpect %+v, got %+v", v.expect, res)
		})
	}
}
