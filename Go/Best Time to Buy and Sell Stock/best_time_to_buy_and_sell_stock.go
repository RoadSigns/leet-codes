package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	maxProfit := 0

	for i, p := range prices {
		for _, q := range prices[i:] {
			if q > p {
				profit := q - p
				if profit > maxProfit {
					maxProfit = profit
				}
			}
		}
	}

	return maxProfit
}
