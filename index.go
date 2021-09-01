package linq

func (q *query[T]) Index(f func(T) bool) int {
	for k, v := range *q.Val {
		if f(v) {
			return k
		}
	}
	return -1
}
