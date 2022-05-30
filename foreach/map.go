package foreach

type MapFunc[T, R any] func(index int, it *T) R

// TODO: implement
//func (c *collection[T]) Map(f MapFunc[T, R]) []R {
//}
