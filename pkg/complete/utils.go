package complete

import (
	"reflect"
	"unsafe"

	"github.com/cornelk/hashmap"
	"golang.org/x/exp/constraints"
)

func GetUnexportedField[E any, R any](entity E, fieldName string) R {
	field := reflect.ValueOf(entity).Elem().FieldByName(fieldName)
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface().(R)
}

func ClearHashmap[K constraints.Ordered, V any](hmap *hashmap.Map[K, V]) {
	keys := make([]K, 0)
	hmap.Range(func(k K, v V) bool {
		keys = append(keys, k)
		return true
	})
	for _, i := range keys {
		hmap.Del(i)
	}
}

func Contains[T constraints.Ordered](arr []T, toCheck T) bool {
	for _, item := range arr {
		if item == toCheck {
			return true
		}
	}

	return false
}

// private fun <E> MutableList<E>.push(element: E) {
//     this.add(element)
// }

// private fun <E> MutableList<E>.pop() : E {
//     val initialSize = this.size
//     val el = this.removeAt(this.size - 1)
//     val afterSize = this.size
//     if (afterSize != (initialSize - 1)) {
//         throw RuntimeException()
//     }
//     return el
// }
