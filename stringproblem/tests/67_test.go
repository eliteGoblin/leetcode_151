package tests

import (
	"fmt"
	"leetcode/stringproblem"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases67 = []struct {
	a      string
	b      string
	expect string
}{
	{
		a:      "11",
		b:      "1",
		expect: "100",
	},
}

func TestAddBinary(t *testing.T) {
	for i, v := range testCases67 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := stringproblem.AddBinary(v.a, v.b)
			assert.Equal(t, v.expect, res)
		})
	}
}
