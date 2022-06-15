package option

import "github.com/versilis/go-util/internal/typeutils"

func Just[T any](val T) Optional[T] {
	if typeutils.IsZero(val) {
		panic("just constructor called with zero value")
	}
	return optionalImpl[T]{val: val}
}

func Try[T any](val T) Optional[T] {
	return optionalImpl[T]{val: val, isEmpty: typeutils.IsZero(val)}
}

func Empty[T any]() Optional[T] {
	return optionalImpl[T]{isEmpty: true}
}
