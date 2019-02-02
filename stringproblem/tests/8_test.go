package tests

import (
	"fmt"
	"leetcode/stringproblem"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases8 = []struct {
	in     string
	expect int
}{
	{
		in:     "1234",
		expect: 1234,
	},
	{
		in:     "-22",
		expect: -22,
	},
}

func TestMyAtoi(t *testing.T) {
	for i, v := range testCases8 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := stringproblem.MyAtoi(v.in)
			assert.Equal(t, v.expect, res)
		})
	}
}
