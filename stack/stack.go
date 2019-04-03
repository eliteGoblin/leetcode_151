package stack

import (
	"errors"
)

type Stack struct {
	curPos int
	stk    []interface{}
}

func NewStack() *Stack {
	return &Stack{
		stk:    make([]interface{}, 2048),
		curPos: 0,
	}
}

func (s *Stack) Top() interface{} {
	if s.curPos == 0 {
		return nil
	}
	return s.stk[s.curPos-1]
}

func (s *Stack) Push(v interface{}) {
	if s.curPos >= len(s.stk) {
		newSpace := make([]interface{}, len(s.stk))
		s.stk = append(s.stk, newSpace...)
	}
	s.stk[s.curPos] = v
	s.curPos++

}

func (s *Stack) Pop() (interface{}, error) {
	if s.curPos == 0 {
		return nil, errors.New("underflow")
	}
	res := s.stk[s.curPos-1]
	s.curPos--
	return res, nil
}

func (s *Stack) Empty() bool {
	return s.curPos == 0
}
