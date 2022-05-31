package apply_test

import (
	"github.com/stretchr/testify/require"
	"go-ahead/apply"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("returns first book whose author contains an r", func(t *testing.T) {
		expected := book{"Charlie"}

		actual := apply.Find(testSet, func(i int, it book) bool {
			return strings.Contains(it.Author, "r")
		})

		require.Equal(t, expected, *actual)
	})

	t.Run("returns book of given author", func(t *testing.T) {
		expected := book{"Erin"}

		actual := apply.Find(testSet, func(i int, it book) bool {
			return it.Author == "Erin"
		})

		require.Equal(t, expected, *actual)
	})

	t.Run("returns no books for unknown author", func(t *testing.T) {
		var expected *book = nil

		actual := apply.Find(testSet, func(i int, it book) bool {
			return it.Author == "Mario"
		})

		require.Equal(t, expected, actual)
	})
}
