package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/c-plus-minus/go-ahead/testkit"
	"github.com/c-plus-minus/go-ahead/testkit/expect"
	"strings"
	"testing"
)

func TestFindIndex(t *testing.T) {
	t.Run("returns index of first book whose author contains an r", func(t *testing.T) {
		wanted := 2

		got := apply.FindIndex(testkit.Set, func(i int, it testkit.Book) bool {
			return strings.Contains(it.Author, "r")
		})

		expect.Equal(t, wanted, got)
	})

	t.Run("returns index of book with given author", func(t *testing.T) {
		wanted := 3

		got := apply.FindIndex(testkit.Set, func(i int, it testkit.Book) bool {
			return it.Author == "Erin"
		})

		expect.Equal(t, wanted, got)
	})

	t.Run("returns negative index for unknown author", func(t *testing.T) {
		wanted := -1

		got := apply.FindIndex(testkit.Set, func(i int, it testkit.Book) bool {
			return it.Author == "Mario"
		})

		expect.Equal(t, wanted, got)
	})
}
