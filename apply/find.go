package apply

func Find[T comparable](slice []T, f func(i int, it T) bool) *T {
	for i, it := range slice {
		if f(i, it) {
			return &it
		}
	}
	return nil
}
