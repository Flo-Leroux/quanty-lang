package list

// List T -
type List[T any] struct {
	typ   string
	items []T
}

func New[T any](typ string) *List[T] {
	return &List[T]{
		items: make([]T, 0),
		typ:   typ,
	}
}

// Type -
func (l *List[T]) Type() string {
	return l.typ
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
