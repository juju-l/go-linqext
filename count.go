package linq

func (q *Query[T]) Count() int {
	return len(*q)
}
