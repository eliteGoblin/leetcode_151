package dp

import (
	"errors"
)

func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := getMaxAreaFromHisto(matrix[0])
	for i := 1; i < len(matrix); i++ {
		for j := range matrix[i] {
			if matrix[i][j] != '0' {
				matrix[i][j] += matrix[i-1][j] - '0'
			}
		}
		maxArea = Max(maxArea, getMaxAreaFromHisto(matrix[i]))
	}
	return maxArea
}

func getMaxAreaFromHisto(histo []byte) int {
	stack := NewStack()
	histo = append(histo, 0)
	maxArea := 0
	for i := 0; i < len(histo); {
		if stack.Empty() || histo[i] > histo[stack.Top()] {
			stack.Push(i)
			i++
			continue
		}
		top := stack.Top()
		stack.Pop()
		if stack.Empty() {
			maxArea = Max(maxArea, int(histo[top]-'0')*i)
		} else {
			maxArea = Max(maxArea, int(histo[top]-'0')*(i-stack.Top()-1))
		}
	}
	return maxArea
}

type stack struct {
	s []int
}

func NewStack() *stack {
	return &stack{make([]int, 0, 4096)}
}

func (s *stack) Top() int {
	return s.s[len(s.s)-1]
}

func (s *stack) Push(v int) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) Empty() bool {
	return len(s.s) == 0
}

func Max(elements ...int) int {
	m := elements[0]
	for i := 1; i < len(elements); i++ {
		if elements[i] > m {
			m = elements[i]
		}
	}
	return m
}
