package linq

func (q *query[T]) Add(v T) *query[T] {
	*q.Val = append(*q.Val, v)
	return q
}
