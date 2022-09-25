package ast

// List T -
type List[T interface{}] struct {
	items []T
}

func NewList[T interface{}]() *List[T] {
	return &List[T]{
		items: make([]T, 0),
	}
}

// Len -
func (l *List[T]) Len() int {
	return len(l.items)
}

// IsEmpty -
func (l *List[T]) IsEmpty() bool {
	return l.Len() == 0
}

// Items -
func (l *List[T]) Items() []T {
	return l.items
}

// Append -
func (l *List[T]) Append(items ...T) {
	l.items = append(l.items, items...)
}
