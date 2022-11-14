package maps

import (
	"fmt"

	"github.com/Flo-Leroux/quanty-lang/pkg/utils/list"

	"golang.org/x/exp/constraints"
)

type Fake struct {
	item bool
}

func (f *Fake) len() string {
	return "FAKE LEN"
}

type Key constraints.Ordered
type Val any

// Map is an ordered map with some extra useful methods
type Map[K Key, V Val] struct {
	entries map[K]Value[K, V]
	keys    *list.List[K]
}

// Value represents the value of a map entry
type Value[K Key, V Val] struct {
	val     V
	element *list.Element[K]
}

// Entry represents a map key value pair entry
type Entry[K Key, V Val] struct {
	Key   K
	Value V
}

// New creates a new Map and returns the Map pointer
func New[K Key, V Val]() *Map[K, V] { return new(Map[K, V]).Init() }

// Init initializes/clears the Map and returns the Map pointer
func (t *Map[K, V]) Init() *Map[K, V] {
	t.entries = make(map[K]Value[K, V])
	if t.keys == nil {
		t.keys = list.New[K]()
	} else {
		t.keys.Init()
	}
	return t
}

// Put enters a new key value pair into the map and returns the Map pointer
func (t *Map[K, V]) Put(key K, value V) *Map[K, V] {
	if v, ok := t.entries[key]; ok {
		t.entries[key] = Value[K, V]{value, v.element}
	} else {
		t.entries[key] = Value[K, V]{value, t.keys.PushBack(key)}
	}
	return t
}

// PutEntries enters all entries into the Map and returns the Map pointer
func (t *Map[K, V]) PutEntries(entries []Entry[K, V]) *Map[K, V] {
	for _, entry := range entries {
		t.Put(entry.Key, entry.Value)
	}
	return t
}

// Delete deletes the key from the map and returns the Map pointer
func (t *Map[K, V]) Delete(key K) *Map[K, V] {
	if _, ok := t.entries[key]; ok {
		t.keys.Remove(t.entries[key].element)
		delete(t.entries, key)
	}
	return t
}

// DeleteAll deletes all the keys from the map and returns the Map pointer
func (t *Map[K, V]) DeleteAll(keys []K) *Map[K, V] {
	for _, k := range keys {
		t.Delete(k)
	}
	return t
}

// Keys creates and returns a slice of all the map keys
func (t *Map[K, V]) Keys() []K {
	keys := make([]K, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		keys[i] = e.Value
	}
	return keys
}

// String returns a string representing the map entries
func (t *Map[K, V]) String() string {
	return fmt.Sprint(t.Entries())
}

// Values creates and returns a slice of all the map values
func (t *Map[K, V]) Values() []V {
	vals := make([]V, t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		vals[i] = t.entries[e.Value].val
	}
	return vals
}

// Entries creates and returns a slice of all the map key value pair entries
func (t *Map[K, V]) Entries() []Entry[K, V] {
	entries := make([]Entry[K, V], t.keys.Len())
	for e, i := t.keys.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		entries[i] = Entry[K, V]{Key: e.Value, Value: t.entries[e.Value].val}
	}
	return entries
}

// Get gets and returns the value for the specified search key
func (t *Map[K, V]) Get(key K) V {
	if val, ok := t.entries[key]; ok {
		return val.val
	}

	var empty V
	return empty
}

// Contains returns true if the map contains the key
func (t *Map[K, V]) Contains(key K) bool {
	_, ok := t.entries[key]
	return ok
}

// ContainsAll returns true if the map contains all the keys
func (t *Map[K, V]) ContainsAll(keys []K) bool {
	for _, key := range keys {
		if !t.Contains(key) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the map contains any of the keys
func (t *Map[K, V]) ContainsAny(keys []K) bool {
	for _, key := range keys {
		if t.Contains(key) {
			return true
		}
	}
	return false
}

// Len returns the number of map entries
func (t *Map[K, V]) Len() int {
	return t.keys.Len()
}

// Head returns the first value of the ordered map
func (t *Map[K, V]) Head() V {
	if t.NotEmpty() {
		return t.entries[t.keys.Front().Value].val
	}

	var empty V
	return empty
}

// Tail returns the last value of the ordered map
func (t *Map[K, V]) Tail() V {
	if t.NotEmpty() {
		return t.entries[t.keys.Back().Value].val
	}

	var empty V
	return empty
}

// Pop deletes the last map entry and returns its value
func (t *Map[K, V]) Pop() V {
	if t.NotEmpty() {
		key := t.keys.Remove(t.keys.Back())
		val, ok := t.entries[key]
		delete(t.entries, key)
		if ok {
			return val.val
		}
	}

	var empty V
	return empty
}

// Pull deletes the first map entry and returns its value
func (t *Map[K, V]) Pull() V {
	if t.NotEmpty() {
		key := t.keys.Remove(t.keys.Front())
		val, ok := t.entries[key]
		delete(t.entries, key)
		if ok {
			return val.val
		}
	}

	var empty V
	return empty
}

// Empty returns true if the Map is empty
func (t *Map[K, V]) Empty() bool {
	return t.keys.Len() == 0
}

// NotEmpty returns true if the Map is not empty
func (t *Map[K, V]) NotEmpty() bool {
	return t.keys.Len() > 0
}
