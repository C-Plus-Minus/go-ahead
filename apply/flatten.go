package apply

import (
	"reflect"
)

func Flatten[T any](slice any) (flat []T) {
	return flatten[T](slice, nil)
}

func flatten[T any](iterable any, remember []T) []T {
	switch reflect.TypeOf(iterable).Kind() {
	case reflect.Slice | reflect.Array:
		v := reflect.ValueOf(iterable)
		for i := 0; i < v.Len(); i++ {
			remember = flatten(v.Index(i).Interface(), remember)
		}
	default:
		remember = append(remember, iterable.(T))
	}
	return remember
}
