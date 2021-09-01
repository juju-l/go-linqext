package linq

func New[T any]() *query[T] {
	return &query[T]{new(List[T]), nil}
}
