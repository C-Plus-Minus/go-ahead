package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/c-plus-minus/go-ahead/testkit"
	"github.com/c-plus-minus/go-ahead/testkit/expect"
	"testing"
)

func TestDo(t *testing.T) {
	t.Run("executes given function with each element of the slice", func(t *testing.T) {
		m := make(map[int]string)

		apply.Do(testkit.Set, func(i int, it testkit.Book) {
			m[len(testkit.Set)-1-i] = it.Author
		})

		expect.Equal(t, testkit.Set[3].Author, m[0])
		expect.Equal(t, testkit.Set[2].Author, m[1])
		expect.Equal(t, testkit.Set[1].Author, m[2])
		expect.Equal(t, testkit.Set[0].Author, m[3])
	})
}
