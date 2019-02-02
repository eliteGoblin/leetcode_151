package stack

// var mpParenthis = map[string]string{
// 	"}": "{",
// 	"]": "[",
// 	")": "(",
// }

// func isValid(s string) bool {
// 	stack := NewStack()
// 	for _, v := range s {
// 		switch v {
// 		case '{':
// 			fallthrough
// 		case '[':
// 			fallthrough
// 		case '(':
// 			stack.Push(string(v))
// 		case '}':
// 			fallthrough
// 		case ']':
// 			fallthrough
// 		case ')':
// 			e, err := stack.Pop()
// 			if err != nil {
// 				return false
// 			}
// 			if e != mpParenthis[string(v)] {
// 				return false
// 			}
// 		}
// 	}
// 	return stack.top == 0
// }
