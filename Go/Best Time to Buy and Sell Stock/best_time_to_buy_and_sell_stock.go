package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	bP := prices[0]
	p := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < bP {
			bP = prices[i]
		} else {
			c := prices[i] - bP
			if c > p {
				p = c
			}
		}
	}
	return p
}
