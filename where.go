package linq

func (q *query[T]) Where(f func(T) bool) *query[T] {
	var list = new(List[T])
	for _, v := range *q.Val {
		if f(v) {
			*list = append(*list, v)
		}
	}
	q.Val = list
	return q
}
