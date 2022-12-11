package reverse_string_test

import (
	"github.com/roadsigns/leet-codes/go/reverse_string"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	got := reverse_string.ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	assert.Equal(t, expected, got)
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse_string.ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	}
}
