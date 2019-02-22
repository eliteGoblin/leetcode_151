package tests

import (
	"fmt"
	"leetcode/dp"
	"testing"
)

func Test85(t *testing.T) {
	fmt.Println(dp.MaximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
	}))
}
