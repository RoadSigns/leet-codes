package best_time_to_buy_and_sell_stock_test

import (
	"github.com/roadsigns/leet-codes/go/best_time_to_buy_and_sell_stock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	expected := 5
	got := best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, expected, got)

	expected = 0
	got = best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, expected, got)
}
