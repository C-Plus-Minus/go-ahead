package foreach

type filterFunc[T any] func(e T) bool

func (c *collection[T]) Filter(f filterFunc[T]) []T {
	var result []T
	for _, it := range c.elements {
		if f(it) {
			result = append(result, it)
		}
	}
	return result
}
