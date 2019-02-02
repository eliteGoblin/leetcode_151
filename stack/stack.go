package stack

import (
	"errors"
)

type Stack struct {
	s []interface{}
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0, 4096)}
}

func (s *Stack) Top() interface{} {
	return s.s[len(s.s)-1]
}

func (s *Stack) Push(v interface{}) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (interface{}, error) {

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *Stack) Empty() bool {
	return len(s.s) == 0
}
