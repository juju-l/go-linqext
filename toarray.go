package linq

func (q *query[T]) ToArray() List[T] {
	return *q.Val
}
