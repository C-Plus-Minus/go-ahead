package apply_test

import (
	"github.com/stretchr/testify/require"
	"go-ahead/apply"
	"strings"
	"testing"
)

func TestFindIndex(t *testing.T) {
	t.Run("returns index of first book whose author contains an r", func(t *testing.T) {
		expected := 2

		actual := apply.FindIndex(testSet, func(i int, it book) bool {
			return strings.Contains(it.Author, "r")
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns index of book with given author", func(t *testing.T) {
		expected := 3

		actual := apply.FindIndex(testSet, func(i int, it book) bool {
			return it.Author == "Erin"
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns negative index for unknown author", func(t *testing.T) {
		expected := -1

		actual := apply.FindIndex(testSet, func(i int, it book) bool {
			return it.Author == "Mario"
		})

		require.Equal(t, expected, actual)
	})
}
