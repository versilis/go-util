package option

type Optional[T any] interface {
	Get() (T, bool)
	Else(other T) T
	ElseGet(supplier func() T) T
	ElseErr(err error) (T, error)
	IsEmpty() bool
	IsPresent() bool
}

type optionalImpl[T any] struct {
	val     T
	isEmpty bool
}

func (o optionalImpl[T]) Get() (T, bool) {
	return o.val, o.IsPresent()
}

func (o optionalImpl[T]) Else(other T) T {
	if o.isEmpty {
		return other
	}
	return o.val
}

func (o optionalImpl[T]) ElseGet(supplier func() T) T {
	if o.isEmpty {
		return supplier()
	}
	return o.val
}

func (o optionalImpl[T]) ElseErr(err error) (T, error) {
	if o.isEmpty {
		return o.val, err
	}
	return o.val, nil
}

func (o optionalImpl[T]) IsEmpty() bool {
	return o.isEmpty
}

func (o optionalImpl[T]) IsPresent() bool {
	return !o.IsEmpty()
}

func Map[A, B any](o Optional[A], transform func(A) B) Optional[B] {
	if value, ok := o.Get(); ok {
		return Just[B](transform(value))
	} else {
		return Empty[B]()
	}
}

func FlatMap[A, B any](o Optional[A], transform func(A) Optional[B]) Optional[B] {
	if value, ok := o.Get(); ok {
		return transform(value)
	} else {
		return Empty[B]()
	}
}
