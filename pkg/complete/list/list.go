package list

func New[T any](equalFn func(a, b T) bool) *List[T] {
	return &List[T]{
		l:       make([]T, 0),
		equalFn: equalFn,
	}
}

type List[T any] struct {
	l       []T
	equalFn func(a, b T) bool
}

func (l *List[T]) Add(e T) {
	l.l = append(l.l, e)
}

func (l *List[T]) AddMany(e []T) {
	l.l = append(l.l, e...)
}

func (l *List[T]) AddList(al *List[T]) {
	l.l = append(l.l, al.ToArray()...)
}

func (l *List[T]) Len() int {
	return len(l.l)
}

func (l *List[T]) ToArray() []T {
	return l.l
}

func (l *List[T]) Filter(fn func(T) bool) *List[T] {
	n := New(l.equalFn)

	for _, i := range l.l {
		if fn(i) {
			n.Add(i)
		}
	}

	return n
}
func (l *List[T]) RemoveAt(idx int) T {

	var e T
	if idx < 0 || idx > len(l.l)-1 {
		return e
	}

	e = l.l[idx]

	l.l = append(l.l[:idx], l.l[idx+1:]...)

	return e
}

func (l *List[T]) Pop() T {
	var e T
	if l.Len() == 0 {
		return e
	}

	e = l.l[l.Len()-1]

	l.l = l.l[:l.Len()-1]

	return e
}

func (l *List[T]) Contains(i T) bool {
	for _, e := range l.l {
		if l.equalFn(e, i) {
			return true
		}
	}

	return false
}
