package greedy

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	start, end := 0, len(height)-1
	maxArea := 0
	for start < end {
		area := getArea(height, start, end)
		if area > maxArea {
			maxArea = area
		}
		if height[start] <= height[end] {
			start++
		} else {
			end--
		}
	}
	return maxArea
}

func getArea(height []int, start, end int) int {
	w := end - start
	h := 0
	if height[start] >= height[end] {
		h = height[end]
	} else {
		h = height[start]
	}
	return w * h
}
