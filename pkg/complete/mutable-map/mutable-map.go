package mutablemap

import (
	"golang.org/x/exp/constraints"
)

func New[K constraints.Ordered, V any]() *MutableMap[K, V] {
	return &MutableMap[K, V]{
		list: make(map[K]V, 0),
	}
}

type MutableMap[K constraints.Ordered, V any] struct {
	list map[K]V
}

func (l *MutableMap[K, V]) Clear() {
	l.list = make(map[K]V, 0)
}

func (l *MutableMap[K, V]) Set(k K, v V) {
	l.list[k] = v
}

func (l *MutableMap[K, V]) Get(k K) V {
	if val, ok := l.list[k]; ok {
		return val
	}

	var v V
	return v
}

func (l *MutableMap[K, V]) Range(fn func(k K, v V) bool) {
	for k, v := range l.list {
		fn(k, v)
	}
}

func (l *MutableMap[K, V]) Contains(i K) bool {
	_, ok := l.list[i]
	return ok
}

func (l *MutableMap[K, V]) ToMap() map[K]V {
	return l.list
}

func (l *MutableMap[K, V]) Len() int {
	return len(l.list)
}
