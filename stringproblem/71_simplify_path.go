package stringproblem

type stack struct {
	top  int
	data []string
}

func NewStack() *stack {
	return &stack{
		top:  0,
		data: make([]string, 4096),
	}
}

func (selfPtr *stack) Push(s string) {
	selfPtr.data[selfPtr.top] = s
	selfPtr.top++
}

func (selfPtr *stack) Pop() {
	if selfPtr.top > 0 {
		selfPtr.top--
	}
}

func SimplifyPath(path string) string {
	stack := NewStack()
	var next string
	for path != "" {
		path, next = getNextPath(path)
		switch next {
		case "":
		case ".":
		case "..":
			stack.Pop()
		default:
			stack.Push(next)
		}
	}
	if 0 == stack.top {
		return "/"
	}
	res := ""
	for i := 0; i < stack.top; i++ {
		res += "/" + stack.data[i]
	}
	return res
}

func getNextPath(path string) (leftPath string, next string) {
	var i int
	for i = 0; i < len(path); i++ {
		if '/' == path[i] {
			break
		}
	}
	if i == len(path) {
		return "", path[:i]
	}
	return path[i+1:], path[:i]
}
