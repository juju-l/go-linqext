package linq

func (q *query) ToArray() List[T] {
	return *q.Val
}
