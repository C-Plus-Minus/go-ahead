package apply

func Do[T comparable](slice []T, f func(i int, it T)) {
	for i, it := range slice {
		f(i, it)
	}
}
