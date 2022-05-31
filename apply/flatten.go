package apply

func Flatten[T any](t []T) (flat []T) {
	// FIXME
	for _, tt := range t {
		var f interface{} = tt // workaround, as t.(type) won't compile
		switch f.(type) {
		case T:
			flat = append(flat, f.(T))
		case []T:
			Flatten(f.([]T))
		}
	}

	return flat
}
