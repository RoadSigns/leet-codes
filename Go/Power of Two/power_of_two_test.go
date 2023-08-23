package power_of_two_test

import (
	"github.com/roadsigns/leet-codes/go/power_of_two"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCalculateIntThatIsPowerOfTwo(t *testing.T) {
	got := power_of_two.IsPowerOfTwo(16)
	assert.Equal(t, true, got)
}

func TestCanCalculateIntThatIsNotPowerOfTwo(t *testing.T) {
	got := power_of_two.IsPowerOfTwo(3)
	assert.Equal(t, false, got)
}

func TestCanCalculateIntThatIsZero(t *testing.T) {
	got := power_of_two.IsPowerOfTwo(0)
	assert.Equal(t, false, got)
}

func TestCanCalculateIntThatIsOne(t *testing.T) {
	got := power_of_two.IsPowerOfTwo(1)
	assert.Equal(t, true, got)
}
