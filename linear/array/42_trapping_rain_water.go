package array

// func Trap(height []int) int {
// 	sum := 0
// 	for i := 0; i < len(height); i++ {
// 		l, r := findLeftRightHighest(height, i)
// 		if l != -1 && r != -1 {
// 			if curCol := min(height[l], height[r]) - height[i]; curCol > 0 {
// 				sum += curCol
// 			}
// 		}
// 	}
// 	return sum
// }

// func findLeftRightHighest(height []int, pos int) (leftHighest int, rightHighest int) {
// 	leftHighest, rightHighest = -1, -1
// 	maxValue := height[pos]
// 	// left
// 	for i := pos - 1; i >= 0; i-- {
// 		if height[i] > maxValue {
// 			leftHighest = i
// 			maxValue = height[i]
// 		}
// 	}
// 	// right
// 	maxValue = height[pos]
// 	for i := pos + 1; i < len(height); i++ {
// 		if height[i] > maxValue {
// 			rightHighest = i
// 			maxValue = height[i]
// 		}
// 	}
// 	return
// }

func Trap(height []int) int {
	sum := 0
	leftHighest := make([]int, len(height))
	rightHighest := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		leftHighest[i] = max(leftHighest[i-1], height[i-1])
	}
	for i := len(height) - 2; i >= 1; i-- {
		rightHighest[i] = max(rightHighest[i+1], height[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		minBound := min(leftHighest[i], rightHighest[i])
		if minBound > height[i] {
			sum += minBound - height[i]
		}
	}
	return sum
}
