package linq

func (q *query[T]) Where(f func(T) bool) *query[T] {
	var list = new(List[T])
	for _, v := range *q.Val {
		// for _, v := range q.ToArray() {
		if f(v) {
			*list = append(*list, v)
			// q = q.Add(v)
		}
	}
	q.Val = list
	return q
}
