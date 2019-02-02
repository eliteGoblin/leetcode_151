package stack

// import "leetcode/utils"

// func longestValidParentheses(s string) int {
// 	stack := NewStack()
// 	res := 0
// 	start := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '(' {
// 			stack.Push(i)
// 		} else {
// 			if stack.Empty() {
// 				start = i + 1
// 			} else {
// 				stack.Pop()
// 				if stack.Empty() {
// 					res = utils.Max(res, i-start+1)
// 				} else {
// 					res = utils.Max(res, i-stack.Top())
// 				}
// 			}
// 		}
// 	}
// 	return res
// }
