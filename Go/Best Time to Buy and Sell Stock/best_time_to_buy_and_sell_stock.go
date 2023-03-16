package best_time_to_buy_and_sell_stock

import "math"

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

func MaxProfitFast(prices []int) int {
	maxProfit := 0
	buy := math.MaxUint32
	for i := 0; i < len(prices); i++ {
		if prices[i] > buy {
			if prices[i]-buy > maxProfit {
				maxProfit = prices[i] - buy
			}
		} else {
			buy = prices[i]
		}
	}
	return maxProfit
}
