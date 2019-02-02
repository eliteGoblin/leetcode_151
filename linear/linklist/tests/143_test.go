package tests

import (
	"fmt"
	"leetcode/linear/linklist"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases143 = []struct {
	in     []int
	k      int
	expect []int
}{
	{
		in:     []int{1, 2, 3, 4},
		expect: []int{1, 4, 2, 3},
	},
}

func TestReorderList(t *testing.T) {
	for i, v := range testCases143 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			lst := linklist.Slice2List(v.in)
			linklist.ReorderList(lst)
			outSlice := linklist.List2Slice(lst)
			assert.Equal(t, v.expect, outSlice, "ecpect %+v, got %+v", v.expect, outSlice)
		})
	}
}
