package rex

import (
	"reflect"
	"sort"
	"strings"
)

type Relation struct {
	relations map[string]*relation
}

func NewRelation() *Relation {
	return &Relation{
		relations: map[string]*relation{},
	}
}

// Equals returns true if compared relations are equal and
// returns false otherwise.
func (r *Relation) Equals(s *Relation) bool {
	if r.len() != s.len() {
		return false
	}
	for k, r := range r.relations {
		s, ok := s.relations[k]
		if !ok {
			return false
		}
		if !r.equals(s) {
			return false
		}
	}
	return true
}

// len returns number of stored relations.
func (r *Relation) len() int {
	return len(r.relations)
}

type relation struct {
	tuples []map[string]any
}

func newRelation() *relation {
	return &relation{
		tuples: []map[string]any{},
	}
}

func (r *relation) equals(s *relation) bool {
	if r.len() != s.len() {
		return false
	}
	for _, t := range r.tuples {
		if !s.hasTuple(t) {
			return false
		}
	}
	return true
}

func (r *relation) hasTuple(tuple map[string]any) bool {
	for _, t := range r.tuples {
		if tupleEquals(t, tuple) {
			return true
		}
	}
	return false
}

func tupleEquals(a, b map[string]any) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		w, ok := b[k]
		if !ok {
			return false
		}
		if !valueEquals(v, w) {
			return false
		}
	}
	return true
}

func valueEquals(u, v any) bool {
	if u, ok := u.(*Relation); ok {
		v, ok := v.(*Relation)
		if !ok || !u.Equals(v) {
			return false
		}
	}
	return reflect.DeepEqual(u, v)

}

// len returns number of stored tuples.
func (r *relation) len() int { return len(r.tuples) }

func keys[T any](m map[string]T) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func key(s []string) string {
	sort.Strings(s)
	b := strings.Builder{}
	for _, s := range s {
		if b.Len() > 0 {
			b.WriteString("|")
		}
		b.WriteString(s)
	}
	return b.String()
}

func (r *Relation) InsertTuple(tuple map[string]any) *Relation {
	k := key(keys(tuple))
	if _, ok := r.relations[k]; !ok {
		r.relations[k] = newRelation()
	}
	s := r.relations[k]
	if !s.hasTuple(tuple) {
		s.tuples = append(s.tuples, tuple)
	}
	return r
}

func (r *Relation) InsertOne(pairs ...func() (string, any)) *Relation {
	m := map[string]any{}
	for _, p := range pairs {
		k, v := p()
		m[k] = v
	}
	r.InsertTuple(m)
	return r
}
