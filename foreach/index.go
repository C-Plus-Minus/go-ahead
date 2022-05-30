package foreach

type FindIndexFunc[T any] func(e T) bool

func (c *collection[T]) FindIndex(f FindIndexFunc[T]) int {
	for i, it := range c.elements {
		if f(it) {
			return i
		}
	}
	return -1
}
