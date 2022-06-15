package typeutils

import "reflect"

func IsZero[T any](v T) bool {
	return reflect.ValueOf(&v).Elem().IsZero()
}
