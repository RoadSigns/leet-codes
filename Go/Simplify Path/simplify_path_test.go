package simplify_path_test

import (
	"github.com/roadsigns/leet_codes/simplify_path"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	t.Run("mulitple directory", func(t *testing.T) {
		got := simplify_path.SimplifyPath("/a/./b/../../c/")
		want := "/c"
		assert.Equal(t, want, got)
	})

	t.Run("back path final directory", func(t *testing.T) {
		got := simplify_path.SimplifyPath("/a//b////c/d//././/..")
		want := "/a/b/c"
		assert.Equal(t, want, got)
	})
}
