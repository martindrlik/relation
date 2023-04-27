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

// Len returns the number of stored tuples.
func (r *Relation) Len() (n int) {
	for _, r := range r.relations {
		n += r.len()
	}
	return
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

func (r *Relation) attributes() map[string]struct{} {
	m := map[string]struct{}{}
	for _, r := range r.relations {
		for _, r := range r.tuples {
			for _, r := range r {
				for k := range r {
					m[k] = struct{}{}
				}
				break
			}
		}
	}
	return m
}

func (r *Relation) key() string {
	a := keys(r.attributes())
	sort.Strings(a)
	k := strings.Builder{}
	k.WriteString("{")
	for _, a := range a {
		k.WriteString(a)
		k.WriteString("|")
	}
	k.WriteString("}")
	return k.String()
}

// len returns number of stored relations.
func (r *Relation) len() int {
	return len(r.relations)
}

type relation struct {
	tuples tuples
}

func newRelation() *relation {
	return &relation{
		tuples: tuples{},
	}
}

func (r *relation) equals(s *relation) bool {
	if r.len() != s.len() {
		return false
	}
	for _, t := range r.tuples {
		for _, t := range t {
			if !s.hasTuple(t) {
				return false
			}
		}
	}
	return true
}

func (r *relation) hasTuple(t tuple) bool {
	k, isPartial := t.key()
	if v, ok := r.tuples[k]; ok {
		if !isPartial {
			return true
		}
		for _, v := range v {
			if tupleEquals(t, v) {
				return true
			}
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

func (r *Relation) InsertTuple(t map[string]any) *Relation {
	t = resolveNestedRelation(t)
	k := key(keys(t))
	if _, ok := r.relations[k]; !ok {
		r.relations[k] = newRelation()
	}
	s := r.relations[k]
	if !s.hasTuple(t) {
		k, _ := tuple(t).key()
		s.tuples[k] = append(s.tuples[k], t)
	}
	return r
}

func resolveNestedRelation(t map[string]any) map[string]any {
	m := map[string]any{}
	for k, v := range t {
		switch x := v.(type) {
		case []map[string]any:
			s := make([]*Relation, len(x))
			for i, x := range x {
				s[i] = NewRelation().InsertTuple(x)
			}
			m[k] = s
		case map[string]any:
			m[k] = NewRelation().InsertTuple(x)
		default:
			m[k] = x
		}
	}
	return m
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

func (r *Relation) InsertMany(pairs ...[]func() (string, any)) *Relation {
	for _, pairs := range pairs {
		r.InsertOne(pairs...)
	}
	return r
}

func One(pairs ...func() (string, any)) []func() (string, any) {
	s := []func() (string, any){}
	for _, p := range pairs {
		s = append(s, p)
	}
	return s
}
