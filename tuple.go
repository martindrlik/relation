package rex

import (
	"sort"

	"golang.org/x/exp/maps"
)

type (
	tuplex struct {
		tuple
		meta
	}

	tuple []any

	meta struct {
		isPossible bool
	}

	tupleMap map[string]any
)

func (t tuple) equal(u tuple) bool {
	for i, v := range t {
		if v != u[i] {
			return false
		}
	}
	return true
}

func (tm tupleMap) ktx() (attrsKey, tuplex) {
	ks := maps.Keys(tm)
	sort.Strings(ks)
	t := make(tuple, len(ks))
	for i, k := range ks {
		v := tm[k]
		if nested, ok := v.(map[string]any); ok {
			t[i] = NewRelation().Insert(nested)
		} else {
			t[i] = v
		}
	}
	return attrs(ks).key(), tuplex{t, meta{}}
}

func (t tuple) toMap(a attrs) map[string]any {
	m := make(map[string]any)
	for i, k := range a {
		m[k] = t[i]
	}
	return m
}
