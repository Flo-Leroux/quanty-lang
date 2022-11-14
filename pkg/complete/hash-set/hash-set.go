package hashset

func New[T any](equalFn func(a, b T) bool) *HashSet[T] {
	return &HashSet[T]{
		set:     make([]T, 0),
		equalFn: equalFn,
	}
}

type HashSet[T any] struct {
	set     []T
	equalFn func(a, b T) bool
}

func (h *HashSet[T]) Add(e T) {
	if !h.Contains(e) {
		h.set = append(h.set, e)
	}
}

func (h *HashSet[T]) Contains(e T) bool {
	contain := false

	for _, s := range h.set {
		if h.equalFn(s, e) {
			contain = true
			break
		}
	}

	return contain
}

func (h *HashSet[T]) Get(idx int) T {
	var e T
	if len(h.set) < idx {
		return e
	}

	return h.set[idx]
}

func (h *HashSet[T]) Clear() {
	h.set = make([]T, 0)
}
