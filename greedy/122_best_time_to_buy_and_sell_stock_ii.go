package greedy

func maxProfitII(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	start, end := 0, 1
	profit := 0
	for start < len(prices) && end < len(prices) {
		cur := prices[end] - prices[start]
		if cur >= 0 {
			profit += cur
			start = end
			end++
		} else {
			start = end
			end++
		}
	}
	return profit
}
