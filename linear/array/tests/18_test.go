package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"
)

var testCases18 = []struct {
	in     []int
	target int
}{
	{
		in:     []int{1, 0, -1, 0, -2, 2},
		target: 0,
	},
}

func TestFourSum(t *testing.T) {
	for _, v := range testCases18 {
		fmt.Println(array.FourSum(v.in, v.target))
	}
}
