package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/c-plus-minus/go-ahead/testkit"
	"github.com/c-plus-minus/go-ahead/testkit/expect"
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("returns first book whose author contains an r", func(t *testing.T) {
		wanted := testkit.Book{Author: "Charlie"}

		got := apply.Find(testkit.Set, func(i int, it testkit.Book) bool {
			return strings.Contains(it.Author, "r")
		})

		expect.Equal(t, wanted, *got)
	})

	t.Run("returns book of given author", func(t *testing.T) {
		wanted := testkit.Book{Author: "Erin"}

		got := apply.Find(testkit.Set, func(i int, it testkit.Book) bool {
			return it.Author == "Erin"
		})

		expect.Equal(t, wanted, *got)
	})

	t.Run("returns no books for unknown author", func(t *testing.T) {
		var wanted *testkit.Book = nil

		got := apply.Find(testkit.Set, func(i int, it testkit.Book) bool {
			return it.Author == "Mario"
		})

		expect.Equal(t, wanted, got)
	})
}
