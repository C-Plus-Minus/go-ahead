package apply

func FindIndex[T comparable](slice []T, f func(i int, it T) bool) int {
	for i, it := range slice {
		if f(i, it) {
			return i
		}
	}
	return -1
}
