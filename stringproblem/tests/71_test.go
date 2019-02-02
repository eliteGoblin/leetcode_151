package tests

import (
	"fmt"
	"leetcode/stringproblem"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases71 = []struct {
	in     string
	expect string
}{
	{
		in:     "/home/",
		expect: "/home",
	},
	{
		in:     "/a/./b/../../c/",
		expect: "/c",
	},
}

func TestSimplifyPath(t *testing.T) {
	for i, v := range testCases71 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := stringproblem.SimplifyPath(v.in)
			assert.Equal(t, v.expect, res)
		})
	}
}
