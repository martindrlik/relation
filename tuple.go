package rex

import (
	"reflect"
)

type T map[string]any

func (t T) Equal(other T) bool {
	if len(t) != len(other) {
		return false
	}
	for k, v := range t {
		if !reflect.DeepEqual(v, other[k]) {
			return false
		}
	}
	return true
}

func (t T) Schema() Schema {
	m := map[string]struct{}{}
	for k := range t {
		m[k] = struct{}{}
	}
	return m
}

func (t T) Projection(projection map[string]struct{}) T {
	c := T{}
	for k, v := range t {
		if _, ok := projection[k]; ok {
			c[k] = v
		}
	}
	return c
}
