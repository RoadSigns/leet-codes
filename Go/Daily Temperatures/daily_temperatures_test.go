package daily_temperatures_test

import (
	"github.com/roadsigns/leet_codes/daily_temperatures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	t.Run("", func(t *testing.T) {
		temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
		got := daily_temperatures.DailyTemperatures(temps)
		want := []int{1, 1, 4, 2, 1, 1, 0, 0}
		for i := 0; i < len(want); i++ {
			assert.Equal(t, want[i], got[i])
		}
	})
}
