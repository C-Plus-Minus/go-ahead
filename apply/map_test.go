package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("transforms book slice to string slice", func(t *testing.T) {
		mapped := apply.Map(testSet, func(i int, it book) string {
			return it.Author
		})

		require.Equal(t, testSet[0].Author, mapped[0])
		require.Equal(t, testSet[1].Author, mapped[1])
		require.Equal(t, testSet[2].Author, mapped[2])
		require.Equal(t, testSet[3].Author, mapped[3])
	})

	t.Run("transforms string slice to book slice", func(t *testing.T) {
		strings := []string{testSet[0].Author, testSet[1].Author}
		mapped := apply.Map(strings, func(i int, it string) book {
			return book{it}
		})

		require.Equal(t, testSet[0], mapped[0])
		require.Equal(t, testSet[1], mapped[1])
	})

	t.Run("extracts inner from outer structs", func(t *testing.T) {
		people := []person{
			{
				Name: name{
					First: "Larry",
					Last:  "Wall",
				},
				Comment: "Godfather of Perl",
			},
			{
				Name: name{
					First: "Dennis",
					Last:  "Ritchie",
				},
				Comment: "Godfather of C",
			},
		}

		mapped := apply.Map(people, func(i int, it person) name {
			return it.Name
		})

		require.Equal(t, []name{people[0].Name, people[1].Name}, mapped)
	})
}
