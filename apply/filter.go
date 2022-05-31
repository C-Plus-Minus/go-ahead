package apply

func Filter[T comparable](slice []T, f func(i int, it T) bool) (filtered []T) {
	for i, it := range slice {
		if f(i, it) {
			filtered = append(filtered, it)
		}
	}
	return filtered
}
