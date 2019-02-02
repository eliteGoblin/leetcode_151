package tests

import (
	"fmt"
	"leetcode/linear/array"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesII(t *testing.T) {
	// 1, 2, 2, 3, 3
	testCase := []int{1, 2, 2, 2, 3, 3}
	res := array.RemoveDuplicatesII(testCase)
	fmt.Println(testCase)
	assert.True(t, res == 5, "%+v", testCase)
}
