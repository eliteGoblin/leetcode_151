package tree

import "errors"

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
	}
}

func (selfPtr *Queue) PushBack(v interface{}) {
	selfPtr.data = append(selfPtr.data, v)
}

func (selfPtr *Queue) Len() int {
	return len(selfPtr.data)
}

func (selfPtr *Queue) PopFront() (v interface{}, err error) {
	if len(selfPtr.data) > 0 {
		res := selfPtr.data[0]
		selfPtr.data = selfPtr.data[1:]
		return res, nil
	}
	return nil, errors.New("underflow")
}

func (selfPtr *Queue) Empty() bool {
	return len(selfPtr.data) == 0
}
