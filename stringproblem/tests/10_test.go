package tests

import (
	"fmt"
	"leetcode/stringproblem"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases10 = []struct {
	s      string
	p      string
	expect bool
}{
	{
		s:      "abc",
		p:      "a.c",
		expect: true,
	},
	{
		s:      "abc",
		p:      "a.",
		expect: false,
	},
	{
		s:      "aa",
		p:      "a*",
		expect: true,
	},
	{
		s:      "ab",
		p:      ".*",
		expect: true,
	},
	{
		s:      "aab",
		p:      "c*a*b",
		expect: true,
	},
}

func TestRE(t *testing.T) {
	for i, v := range testCases10 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := stringproblem.IsMatch(v.s, v.p)
			assert.Equal(t, v.expect, res)
		})
	}
}
