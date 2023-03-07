package schema

import (
	"errors"
	"maps"
	"reflect"
)

var (
	ErrEmpty    = errors.New("empty schema")
	ErrMismatch = errors.New("schema mismatch")
)

type Schema map[string]reflect.Type

func FromTuple(t map[string]any) Schema {
	s := make(Schema)
	for k, v := range t {
		s[k] = reflect.TypeOf(v)
	}
	return s
}

// Equal reports whether two schemas contain the same key/value pairs.
// Values are compared using ==.
func (s Schema) Equal(t Schema) bool { return maps.Equal(s, t) }

// Intersection returns a new schema with the common key/value pairs of s and t.
func (s Schema) Intersection(t Schema) Schema {
	w := make(Schema)
	for a, st := range s {
		if tt, ok := t[a]; ok && st == tt {
			w[a] = st
		}
	}
	return w
}

// Has reports whether s contains the key k.
func (s Schema) Has(k string) bool { _, ok := s[k]; return ok }
