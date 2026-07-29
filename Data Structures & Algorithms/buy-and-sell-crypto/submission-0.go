func maxProfit(prices []int) int {
	myProfit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		currProfit := prices[i] - buy
		if myProfit > currProfit {
			buy = min(buy, prices[i])
		} else {
			myProfit = currProfit
		}
	}

	return myProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}