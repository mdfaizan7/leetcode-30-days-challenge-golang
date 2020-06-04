package packages

// MaxProfit is the main func for this problem
func MaxProfit(prices []int) int {
	profit := 0
	for i := range prices[:len(prices)-1] {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}

	return profit
}
