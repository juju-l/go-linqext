package linq

type query[T any] struct {
	Val *List[T]
	Err error
}
