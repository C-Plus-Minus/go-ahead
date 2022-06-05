package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/c-plus-minus/go-ahead/testkit"
	"github.com/c-plus-minus/go-ahead/testkit/expect"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Run("returns multiple books whose authors end with an e", func(t *testing.T) {
		wanted := []testkit.Book{{"Alice"}, {"Charlie"}}

		got := apply.Filter(testkit.Set, func(i int, it testkit.Book) bool {
			return strings.HasSuffix(it.Author, "e")
		})

		expect.Equal(t, wanted, append(got))
	})

	t.Run("returns the book whose author does not contain an i", func(t *testing.T) {
		wanted := []testkit.Book{{"Bob"}}

		got := apply.Filter(testkit.Set, func(i int, it testkit.Book) bool {
			return !strings.Contains(it.Author, "i")
		})

		expect.Equal(t, wanted, got)
	})

	t.Run("returns no books for authors longer than 10 characters", func(t *testing.T) {
		var wanted []testkit.Book = nil

		got := apply.Filter(testkit.Set, func(i int, it testkit.Book) bool {
			return len(it.Author) > 10
		})

		expect.Equal(t, wanted, got)
	})
}
