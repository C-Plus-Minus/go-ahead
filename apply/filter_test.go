package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("returns multiple books whose authors end with an e", func(t *testing.T) {
		expected := []book{{"Alice"}, {"Charlie"}}

		actual := apply.Filter(testSet, func(i int, it book) bool {
			return strings.HasSuffix(it.Author, "e")
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns the book whose author does not contain an i", func(t *testing.T) {
		expected := []book{{"Bob"}}

		actual := apply.Filter(testSet, func(i int, it book) bool {
			return !strings.Contains(it.Author, "i")
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns no books for authors longer than 10 characters", func(t *testing.T) {
		var expected []book = nil

		actual := apply.Filter(testSet, func(i int, it book) bool {
			return len(it.Author) > 10
		})

		require.Equal(t, expected, actual)
	})
}
