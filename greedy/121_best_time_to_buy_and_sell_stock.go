package greedy

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy, sell := 0, 0
	max := 0
	for buy < len(prices) && sell < len(prices) {
		if max < prices[sell]-prices[buy] {
			max = prices[sell] - prices[buy]
		} else if prices[sell] < prices[buy] {
			buy = sell
		}
		sell++
	}
	return max
}
