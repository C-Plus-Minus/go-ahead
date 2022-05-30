package foreach

type DoFunc[T any] func(index int, it T)

func (c *collection[T]) Do(f DoFunc[T]) {
	for i, it := range c.elements {
		f(i, it) // False positive: Intellij bug error-marks this falsely
	}
}
