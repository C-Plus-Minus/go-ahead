package apply

func Map[T, R any](slice []T, f func(i int, it T) R) (mapped []R) {
	for i, it := range slice {
		mapped = append(mapped, f(i, it))
	}
	return mapped
}
