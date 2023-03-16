package rex

import (
	"bytes"
	"sort"
)

// R is a set of relations.
type R map[string]Relation

func (r R) attributes() []string {
	m := map[string]struct{}{}
	for _, r := range r {
		for _, a := range r.attributes {
			m[a] = struct{}{}
		}
	}
	o := make([]string, 0, len(m))
	for k := range m {
		o = append(o, k)
	}
	sort.Strings(o)
	return o
}

func (r R) keyOrder() []string {
	o := make([]string, 0, len(r))
	for k := range r {
		o = append(o, k)
	}
	sort.Strings(o)
	return o
}

// Relation is a set of tuples.
type Relation struct {
	attributes []string
	tuples     [][]any
}

func (r Relation) key() string {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() { bufPool.Put(b) }()

	for _, k := range r.attributes {
		if b.Len() > 0 {
			b.WriteRune('|')
		}
		b.WriteString(k)
	}
	return b.String()
}

func (r Relation) attri() map[string]int {
	m := map[string]int{}
	for i, a := range r.attributes {
		m[a] = i
	}
	return m
}
