package rex

import (
	"reflect"
	"sort"
)

type Tuple map[string]any

type tuplex struct {
	Tuple
	metadata
}

type metadata struct {
	isPossible bool
}

func (t Tuple) equal(u Tuple) bool {
	if len(t) != len(u) {
		return false
	}
	for k, tv := range t {
		uv, ok := u[k]
		if !ok {
			return false
		}
		if !equal(tv, uv) {
			return false
		}
	}
	return true
}

func equal(a, b any) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	if _, ok := a.(*Relation); ok {
		panic("not implemented")
	}
	if !reflect.DeepEqual(a, b) {
		return false
	}
	return true
}

func (t Tuple) attrs() attrs {
	a := make([]string, 0, len(t))
	for k := range t {
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}

func (t Tuple) combine(u Tuple) Tuple {
	r := make(Tuple)
	for _, src := range []Tuple{t, u} {
		for k, v := range src {
			r[k] = v
		}
	}
	return r
}
