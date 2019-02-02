package tests

import (
	"fmt"
	"leetcode/stringproblem"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases44 = []struct {
	s      string
	p      string
	expect bool
}{
	{
		s:      "ab",
		p:      "a?",
		expect: true,
	},
	{
		s:      "ab",
		p:      "??",
		expect: true,
	},
	{
		s:      "ab",
		p:      "???",
		expect: false,
	},
	{
		s:      "adceb",
		p:      "*a*b",
		expect: true,
	},
	{
		s:      "aa",
		p:      "*",
		expect: true,
	},
}

func TestWildcardMatch(t *testing.T) {
	for i, v := range testCases44 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := stringproblem.IsMatchWildcard(v.s, v.p)
			assert.Equal(t, v.expect, res)
		})
	}
}
