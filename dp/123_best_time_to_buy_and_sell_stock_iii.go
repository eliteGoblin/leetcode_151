package dp

// 见 grandyang 解释
func maxProfitDP(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	localMax := make([][]int, len(prices))
	for i := 0; i < len(localMax); i++ {
		localMax[i] = make([]int, 3)
	}
	globalMax := make([][]int, len(prices))
	for i := 0; i < len(globalMax); i++ {
		globalMax[i] = make([]int, 3)
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= 2; j++ {
			diff := prices[i] - prices[i-1]
			localMax[i][j] = Max(globalMax[i-1][j-1]+Max(diff, 0), localMax[i-1][j]+diff)
			globalMax[i][j] = Max(globalMax[i-1][j], localMax[i][j])
		}
	}
	return globalMax[len(prices)-1][2]
}

func maxProfitNaive(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		curProfit := findMaxProfixSingleBuy(prices[:i]) + findMaxProfixSingleBuy(prices[i:])
		if curProfit > max {
			max = curProfit
		}
	}
	return max
}

func findMaxProfixSingleBuy(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minValueFromLeft := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minValueFromLeft {
			minValueFromLeft = prices[i]
			continue
		}
		if prices[i]-minValueFromLeft > maxProfit {
			maxProfit = prices[i] - minValueFromLeft
		}
	}
	return maxProfit
}
