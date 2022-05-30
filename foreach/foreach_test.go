package foreach_test

import (
	"github.com/stretchr/testify/require"
	"go-ahead/foreach"
	"strings"
	"testing"
)

type book struct {
	Author string
}

var testSet = []book{
	{"Alice"},
	{"Bob"},
	{"Charlie"},
	{"Erin"},
}

func TestOf_Filter(t *testing.T) {
	t.Run("returns multiple books whose authors end with an e", func(t *testing.T) {
		expected := []book{{"Alice"}, {"Charlie"}}

		actual := foreach.Of(testSet).Filter(func(it book) bool {
			return strings.HasSuffix(it.Author, "e")
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns the book whose author does not contain an i", func(t *testing.T) {
		expected := []book{{"Bob"}}

		actual := foreach.Of(testSet).Filter(func(it book) bool {
			return !strings.Contains(it.Author, "i")
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns no books for authors longer than 10 characters", func(t *testing.T) {
		var expected []book = nil

		actual := foreach.Of(testSet).Filter(func(it book) bool {
			return len(it.Author) > 10
		})

		require.Equal(t, expected, actual)
	})
}

func TestOf_Find(t *testing.T) {
	t.Run("returns first book whose author contains an r", func(t *testing.T) {
		expected := book{"Charlie"}

		actual := foreach.Of(testSet).Find(func(it book) bool {
			return strings.Contains(it.Author, "r")
		})

		require.Equal(t, expected, *actual)
	})

	t.Run("returns book of given author", func(t *testing.T) {
		expected := book{"Erin"}

		actual := foreach.Of(testSet).Find(func(it book) bool {
			return it.Author == "Erin"
		})

		require.Equal(t, expected, *actual)
	})

	t.Run("returns no books for unknown author", func(t *testing.T) {
		var expected *book = nil

		actual := foreach.Of(testSet).Find(func(it book) bool {
			return it.Author == "Mario"
		})

		require.Equal(t, expected, actual)
	})
}

func TestOf_FindIndex(t *testing.T) {
	t.Run("returns index of first book whose author contains an r", func(t *testing.T) {
		expected := 2

		actual := foreach.Of(testSet).FindIndex(func(it book) bool {
			return strings.Contains(it.Author, "r")
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns index of book with given author", func(t *testing.T) {
		expected := 3

		actual := foreach.Of(testSet).FindIndex(func(it book) bool {
			return it.Author == "Erin"
		})

		require.Equal(t, expected, actual)
	})

	t.Run("returns negative index for unknown author", func(t *testing.T) {
		expected := -1

		actual := foreach.Of(testSet).FindIndex(func(it book) bool {
			return it.Author == "Mario"
		})

		require.Equal(t, expected, actual)
	})
}

func TestOf_Do(t *testing.T) {
	t.Run("applies function to each element", func(t *testing.T) {
		m := make(map[int]string)

		foreach.Of(testSet).Do(func(i int, it book) {
			m[len(testSet)-1-i] = it.Author
		})

		require.Equal(t, testSet[3].Author, m[0])
		require.Equal(t, testSet[2].Author, m[1])
		require.Equal(t, testSet[1].Author, m[2])
		require.Equal(t, testSet[0].Author, m[3])
	})
}

func TestOf_Map(t *testing.T) {
	t.Run("returns transformed set", func(t *testing.T) {

	})
}
