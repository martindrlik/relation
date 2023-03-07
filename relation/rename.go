package relation

import (
	"maps"

	"github.com/martindrlik/rex/schema"
)

// Rename1 returns a new relation with the attribute old renamed to new.
func (r *Relation) Rename1(old, new string) (*Relation, error) {
	if !r.Schema.Has(old) || r.Schema.Has(new) {
		return nil, schema.ErrMismatch
	}

	schema := make(schema.Schema, len(r.Schema))
	for k, v := range r.Schema {
		schema[k] = v
	}
	schema[new] = schema[old]
	delete(schema, old)

	w, _ := New(schema)
	for t := range r.List() {
		wt := maps.Clone(t)
		wt[new] = wt[old]
		delete(wt, old)
		w.TupleSet.Add(wt)
	}
	return w, nil
}
