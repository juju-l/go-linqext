package linq

func (q *Query[T]) Select(f func(T) bool) *Query[T] {
	var t = new(Query[T])
	for _, v := range *q {
		if f(v) {
			*t = append(*t, v)
			// t = t.Add(v)
		}
	}
	return t
}
