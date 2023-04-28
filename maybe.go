package opt

type Maybe[T any] struct {
	Value  T
	Exists bool
}

func Some[T any](value T) Maybe[T] {
	return Maybe[T]{value, true}
}
func No[T any]() Maybe[T] {
	return Maybe[T]{Exists: false}
}
func First[T any](value []T) Maybe[T] {
	if value == nil {
		return No[T]()
	} else {
		return Some(value[0])
	}
}
func (n Maybe[T]) orElse(value T) T {
	if n.Exists {
		return n.Value
	}
	return value
}
func (n Maybe[T]) orPanic() T {
	if !n.Exists {
		panic("attempted to access invalid Maybe")
	}
	return n.Value
}
