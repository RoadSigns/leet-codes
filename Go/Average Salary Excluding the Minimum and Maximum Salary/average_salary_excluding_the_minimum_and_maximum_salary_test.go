package average_salary_excluding_the_minimum_and_maximum_salary_test

import (
	"github.com/roadsigns/leet-codes/go/average_salary_excluding_the_minimum_and_maximum_salary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverage(t *testing.T) {
	assert.Equal(t, 2000.0, average_salary_excluding_the_minimum_and_maximum_salary.Average([]int{1000, 2000, 3000}))
	assert.Equal(t, 2500.0, average_salary_excluding_the_minimum_and_maximum_salary.Average([]int{4000, 3000, 1000, 2000}))
}

func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		average_salary_excluding_the_minimum_and_maximum_salary.Average([]int{4000, 3000, 1000, 2000})
	}
}
