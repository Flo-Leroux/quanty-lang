package mutablelist

func New[T any](equalFn func(a, b T) bool) *MutableList[T] {
	return &MutableList[T]{
		list:    make([]T, 0),
		equalFn: equalFn,
	}
}

type MutableList[T any] struct {
	list    []T
	equalFn func(a, b T) bool
}

func (l *MutableList[T]) Len() int {
	return len(l.list)
}

func (l *MutableList[T]) Add(e T) {
	l.list = append(l.list, e)
}

func (l *MutableList[T]) AddList(el *MutableList[T]) {
	l.list = append(l.list, el.ToArray()...)
}

func (l *MutableList[T]) Clear() {
	l.list = make([]T, 0)
}

func (l *MutableList[T]) ToArray() []T {
	return l.list
}

func (l *MutableList[T]) Get(idx int) T {
	var e T
	if len(l.list)-1 < idx {
		return e
	}

	return l.list[idx]
}

func (l *MutableList[T]) Slice(start, stop int) *MutableList[T] {
	n := New(l.equalFn)
	if start < 0 || stop > len(l.list) {
		return n
	}

	for i := start; i <= stop; i++ {
		n.Add(l.list[i])
	}

	return n
}

func (l *MutableList[T]) Pop() T {
	var e T
	if l.Len() == 0 {
		return e
	}

	e = l.list[l.Len()-1]

	l.list = l.list[:l.Len()-1]

	return e
}

func (l *MutableList[T]) IndexOf(e T) int {

	for i, t := range l.list {
		if l.equalFn(t, e) {
			return i
		}
	}

	return -1
}
