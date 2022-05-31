package apply_test

import (
	"github.com/stretchr/testify/require"
	"go-ahead/apply"
	"testing"
)

func TestFlatten(t *testing.T) {
	t.Run("transforms book slice to string slice", func(t *testing.T) {
		test := [][][]string{
			{
				{"bla"},
				{"blubb"},
			},
			{
				{"tralala"},
			},
		}
		x := apply.Flatten(test)

		require.Equal(t, []string{"bla, blubb"}, x)
	})

	/*
		type batch struct {
			ID    int
			Books []book
		}

		t.Run("transforms complex slice to flat book slice", func(t *testing.T) {
			batches := []batch{
				{
					ID:    0,
					Books: testSet[0:2],
				},
				{
					ID:    1,
					Books: testSet[2:],
				},
			}

			mapped := apply.Map(batches, func(i int, it batch) []book {
				m := make(map[int][]book)
				m[it.ID] = it.Books
				for _, b := range it.Books {
					books = append(books, b)
				}
				return books
			})

			require.Equal(t, testSet, mapped)
		})
	*/
}
