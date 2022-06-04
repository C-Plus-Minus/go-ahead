package apply_test

import (
	"github.com/stretchr/testify/require"
	"go-ahead/apply"
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

		require.Equal(t, []string{"Sorry Mario", "but", "our", "princess is in another castle"}, flattened)
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

		require.Equal(t, []any{10, true, false, "string", "array", 4.2}, flattened)
	})
}
