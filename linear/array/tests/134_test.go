package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases134 = []struct {
	gas    []int
	cost   []int
	expect int
}{
	{
		gas:    []int{1, 2, 3, 4, 5},
		cost:   []int{3, 4, 5, 1, 2},
		expect: 3,
	},
	{
		gas:    []int{2, 3, 4},
		cost:   []int{3, 4, 3},
		expect: -1,
	},
}

func TestIsCanCompleteCircuit(t *testing.T) {
	for i, tt := range testCases134 {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := array.IsCanCompleteCircuit(tt.gas, tt.cost)
			assert.Equal(t, tt.expect, res)
		})
	}
}
