package stack

// import "leetcode/utils"

// func largestRectangleArea(heights []int) int {
// 	stack := NewStack()
// 	heights = append(heights, 0)
// 	result := 0
// 	for i := 0; i < len(heights); {
// 		if stack.Empty() || heights[stack.Top()] < heights[i] {
// 			stack.Push(i)
// 			i++
// 		} else {
// 			top := stack.Top()
// 			stack.Pop()
// 			if stack.Empty() {
// 				result = utils.Max(result, heights[top]*i)
// 			} else {
// 				result = utils.Max(result, heights[top]*(i-stack.Top()-1))
// 			}
// 		}
// 	}
// 	return result
// }
