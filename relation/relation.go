package relation

import (
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

// Relation is a set of tuples with a common schema.
type Relation struct {
	schema.Schema
	tuple.TupleSet
}

// New creates a new relation with the given schema.
func New(s schema.Schema) (*Relation, error) {
	if len(s) == 0 {
		return nil, schema.ErrEmpty
	}
	return &Relation{Schema: s}, nil
}

// Add adds a tuple to the relation. If the tuple is already in the relation, it does nothing.
// If the tuple has a different schema than the relation, it returns ErrSchemaMismatch.
func (r *Relation) Add(t tuple.Tuple) error {
	if !r.Schema.Equal(schema.FromTuple(t)) {
		return schema.ErrMismatch
	}
	if !r.TupleSet.Has(t) {
		r.TupleSet.Add(t)
	}
	return nil
}
