package power_of_three_test

import (
	"github.com/roadsigns/leet-codes/go/power_of_three"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCalculateIntThatIsPowerOfThree(t *testing.T) {
	got := power_of_three.IsPowerOfThree(27)
	assert.Equal(t, true, got)
}

func TestCanCalculateIntThatIsNotPowerOfThree(t *testing.T) {
	got := power_of_three.IsPowerOfThree(28)
	assert.Equal(t, false, got)
}

func TestCanCalculateIntThatIsZero(t *testing.T) {
	got := power_of_three.IsPowerOfThree(0)
	assert.Equal(t, false, got)
}
