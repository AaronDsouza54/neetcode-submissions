func maxProfit(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		for j:= 0; j < i; j++ {
			temp := prices[i] - prices[j]
			if temp > profit {
				profit = temp
			}
		}
	}

	return profit
}
