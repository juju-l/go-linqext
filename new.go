package linq

func New[T any]() {
	return &query[T]{new(List[T]), nil}
}
