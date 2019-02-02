package tests

import (
	"fmt"
	"leetcode/stringproblem"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases5 = []struct {
	in     string
	expect string
}{
	// {
	// 	in:     "abab",
	// 	expect: "aba",
	// },
	// {
	// 	in:     "babad",
	// 	expect: "bab",
	// },
	{
		in:     "abcda",
		expect: "a",
	},
}

func TestLongestPalindrome(t *testing.T) {
	for i, v := range testCases5 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := stringproblem.LongestPalindromeMemo(v.in)
			assert.Equal(t, v.expect, res)
		})
	}
}
