package linq

type Query[T any] struct {
	Val *List[T]
	Err error
}
