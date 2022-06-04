package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDo(t *testing.T) {
	t.Run("executes given function with each element of the slice", func(t *testing.T) {
		m := make(map[int]string)

		apply.Do(testSet, func(i int, it book) {
			m[len(testSet)-1-i] = it.Author
		})

		require.Equal(t, testSet[3].Author, m[0])
		require.Equal(t, testSet[2].Author, m[1])
		require.Equal(t, testSet[1].Author, m[2])
		require.Equal(t, testSet[0].Author, m[3])
	})
}
