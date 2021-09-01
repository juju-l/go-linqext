package linq

func (q *query[T]) Count() int {
	return len(*q.Val)
}
