package mutableset

func New[V any](equalFn func(a, b V) bool) *MutableSet[V] {
	return &MutableSet[V]{
		list:    make([]V, 0),
		equalFn: equalFn,
	}
}

type MutableSet[V any] struct {
	list    []V
	equalFn func(a, b V) bool
}

func (l *MutableSet[V]) Add(e V) {
	l.list = append(l.list, e)
}

func (l *MutableSet[V]) Clear() {
	l.list = make([]V, 0)
}

func (l *MutableSet[V]) Contains(e V) bool {
	for _, i := range l.list {
		if l.equalFn(i, e) {
			return true
		}
	}

	return false
}

func (l *MutableSet[V]) ToArray() []V {
	return l.list
}
