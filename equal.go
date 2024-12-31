package deep

import (
	"maps"
	"reflect"
	"slices"
	"unsafe"
)

func EqualSlices[S interface{ ~[]E }, E comparable](s1 S, s2 S) bool {
	return slices.EqualFunc(s1, s2, func(v1, v2 E) bool {
		return equalValues(v1, v2)
	})
}

func EqualMaps[M1 interface{ ~map[K]V }, M2 interface{ ~map[K]V }, K comparable, V comparable](m1 M1, m2 M2) bool {
	return maps.EqualFunc(m1, m2, func(v1, v2 V) bool {
		return equalValues(v1, v2)
	})
}

func equalValues[V comparable](v1, v2 V) bool {
	val1 := reflect.ValueOf(v1)
	val2 := reflect.ValueOf(v2)

	switch val1.Kind() {
	case reflect.Slice, reflect.Array:
		if val1.Len() == 0 && val2.Len() != 0 || val1.Len() != 0 && val2.Len() == 0 {
			return false
		}
		if val1.Len() != val2.Len() {
			return false
		}
		u1 := val1.Slice(0, 0).UnsafePointer()
		u2 := val2.Slice(0, 0).UnsafePointer()

		return EqualSlices(
			unsafe.Slice(&u1, val1.Len()),
			unsafe.Slice(&u2, val2.Len()),
		)
	case reflect.Map:
		iter1 := val1.MapRange()
		for iter1.Next() {
			k, v1 := iter1.Key(), iter1.Value()
			if v2 := val2.MapIndex(k); v2.IsZero() && !v1.IsZero() || !equalValues(v1.Interface(), v2.Interface()) {
				return false
			}
		}

		return true
	}

	return v1 == v2
}
