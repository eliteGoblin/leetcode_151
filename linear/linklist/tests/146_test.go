package tests

import (
	"leetcode/linear/linklist"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := linklist.Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
}
