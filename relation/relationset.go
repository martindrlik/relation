package relation

import "github.com/martindrlik/rex/schema"

type RelationSet []*Relation

// Add adds a relation to the set.
func (rs *RelationSet) Add(r *Relation) {
	*rs = append(*rs, r)
}

// Relation returns the relation with the given schema and true if it is in the set.
func (rs *RelationSet) Relation(schema schema.Schema) (*Relation, bool) {
	for _, r := range *rs {
		if r.Schema.Equal(schema) {
			return r, true
		}
	}
	return nil, false
}
