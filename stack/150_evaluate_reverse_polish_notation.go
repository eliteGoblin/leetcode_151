package stack

// import (
// 	"fmt"
// 	"strconv"
// )

// func evalRPN(tokens []string) int {
// 	stack := NewStack()
// 	for _, v := range tokens {
// 		switch v {
// 		case "+":
// 			fallthrough
// 		case "-":
// 			fallthrough
// 		case "*":
// 			fallthrough
// 		case "/":
// 			res := doOperation(stack, v)
// 			stack.Push(res)
// 		default:
// 			stack.Push(v)
// 		}
// 	}
// 	res, _ := strconv.Atoi(stack.Top().(string))
// 	return res
// }

// func doOperation(stack *stack, token string) string {
// 	iOp2, _ := stack.Pop()
// 	iOp1, _ := stack.Pop()
// 	sOp1, _ := strconv.Atoi(iOp1.(string))
// 	sOp2, _ := strconv.Atoi(iOp2.(string))
// 	res := 0
// 	switch token {
// 	case "+":
// 		res = sOp1 + sOp2
// 	case "-":
// 		res = sOp1 - sOp2
// 	case "*":
// 		res = sOp1 * sOp2
// 	case "/":
// 		res = sOp1 / sOp2
// 	}
// 	return fmt.Sprintf("%d", res)
// }
