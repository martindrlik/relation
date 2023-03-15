package rex

import (
	"sort"
	"strings"
)

// Relation is a set of tuples.
type Relation struct {
	attributes []string
	tuples     [][]any
}

// R is a set of relations.
type R map[string]Relation

func newRelation(m map[string]any) Relation {
	r := Relation{}
	r.attributes = make([]string, 0, len(m))
	r.tuples = make([][]any, 1)
	r.tuples[0] = make([]any, len(m))
	for k := range m {
		r.attributes = append(r.attributes, k)
	}
	sort.Strings(r.attributes)
	for i, k := range r.attributes {
		r.tuples[0][i] = m[k]
	}
	return r
}

func (r Relation) key() string {
	var sb strings.Builder
	for _, k := range r.attributes {
		if sb.Len() > 0 {
			sb.WriteRune('|')
		}
		sb.WriteString(k)
	}
	return sb.String()
}

func (r Relation) attri() map[string]int {
	m := map[string]int{}
	for i, a := range r.attributes {
		m[a] = i
	}
	return m
}

func (r R) keyOrder() []string {
	o := make([]string, 0, len(r))
	for k := range r {
		o = append(o, k)
	}
	sort.Strings(o)
	return o
}
