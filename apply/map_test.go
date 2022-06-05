package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/c-plus-minus/go-ahead/testkit"
	"github.com/c-plus-minus/go-ahead/testkit/expect"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("transforms book slice to string slice", func(t *testing.T) {
		mapped := apply.Map(testkit.Set, func(i int, it testkit.Book) string {
			return it.Author
		})

		expect.Equal(t, testkit.Set[0].Author, mapped[0])
		expect.Equal(t, testkit.Set[1].Author, mapped[1])
		expect.Equal(t, testkit.Set[2].Author, mapped[2])
		expect.Equal(t, testkit.Set[3].Author, mapped[3])
	})

	t.Run("transforms string slice to book slice", func(t *testing.T) {
		strings := []string{testkit.Set[0].Author, testkit.Set[1].Author}
		mapped := apply.Map(strings, func(i int, it string) testkit.Book {
			return testkit.Book{Author: it}
		})

		expect.Equal(t, testkit.Set[0], mapped[0])
		expect.Equal(t, testkit.Set[1], mapped[1])
	})

	t.Run("extracts inner from outer structs", func(t *testing.T) {
		people := []testkit.Person{
			{
				Name: testkit.Name{
					First: "Larry",
					Last:  "Wall",
				},
				Comment: "Godfather of Perl",
			},
			{
				Name: testkit.Name{
					First: "Dennis",
					Last:  "Ritchie",
				},
				Comment: "Godfather of C",
			},
		}

		mapped := apply.Map(people, func(i int, it testkit.Person) testkit.Name {
			return it.Name
		})

		expect.Equal(t, []testkit.Name{people[0].Name, people[1].Name}, mapped)
	})
}
