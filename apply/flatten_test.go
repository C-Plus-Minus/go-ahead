package apply_test

import (
	"github.com/c-plus-minus/go-ahead/apply"
	"github.com/c-plus-minus/go-ahead/testkit/expect"
	"testing"
)

func TestFlatten(t *testing.T) {
	t.Run("transforms homogeneous nested slice to string slice", func(t *testing.T) {
		test := [][][]string{
			{
				{"Sorry Mario"},
				{"but", "our"},
			},
			{
				{"princess is in another castle"},
			},
		}

		flattened := apply.Flatten[string](test)

		expect.Equal(t, []string{"Sorry Mario", "but", "our", "princess is in another castle"}, flattened)
	})

	t.Run("transforms heterogeneous slice to string slice", func(t *testing.T) {
		test := [][]any{
			{
				10,
				[]bool{true, false},
			},
			{
				"string",
				[]string{"array"},
				[][]float64{{4.2}},
			},
		}

		flattened := apply.Flatten[any](test)

		expect.Equal(t, []any{10, true, false, "string", "array", 4.2}, flattened)
	})
}
