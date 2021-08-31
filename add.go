package linq

func (q *Query[T]) Add(v T) *Query[T] {
	*q = append(*q, v)
	return q
}
