package foreach

func Of[T any](e []T) *collection[T] {
	return &collection[T]{e}
}

type collection[T any] struct {
	elements []T
}
