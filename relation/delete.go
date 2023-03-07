package relation

import "github.com/martindrlik/rex/schema"

func (r *Relation) Delete(u map[string]any) {
	us := schema.FromTuple(u)
	if r.Schema.Equal(us) {
		r.TupleSet.Delete(u)
	}
}
