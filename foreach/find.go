package foreach

type FindFunc[T any] func(e T) bool

func (c *collection[T]) Find(f FindFunc[T]) *T {
	for _, it := range c.elements {
		if f(it) {
			return &it
		}
	}
	return nil
}
