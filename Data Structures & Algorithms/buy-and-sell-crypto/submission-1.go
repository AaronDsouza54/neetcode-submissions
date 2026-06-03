func maxProfit(prices []int) int {
	profit := 0

	min_value_idx := 0

	for i := 1; i < len(prices); i++ {
		for j:= i - 1; j < i; j++ {
			if prices[j] < prices[min_value_idx] {
				min_value_idx = j
			}

			temp := prices[i] - prices[min_value_idx]
			if temp > profit {
				profit = temp
			}
		}
	}

	return profit
}
