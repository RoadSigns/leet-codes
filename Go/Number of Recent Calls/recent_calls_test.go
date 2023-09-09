package recent_calls_test

import (
	"github.com/roadsigns/leet_codes/recent_calls"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecentCounter_Ping(t *testing.T) {
	counter := recent_calls.Constructor()

	got := counter.Ping(1)
	want := 1
	assert.Equal(t, want, got)

	got = counter.Ping(100)
	want = 2
	assert.Equal(t, want, got)

	got = counter.Ping(3001)
	want = 3
	assert.Equal(t, want, got)

	got = counter.Ping(3002)
	want = 3
	assert.Equal(t, want, got)
}
