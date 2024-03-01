package relation

import "github.com/martindrlik/rex/schema"

func (r *Relation) Project(s schema.Schema) *Relation {
	nr := &Relation{s: s}
	for _, t := range r.t {
		nr.Append(t.Project(s))
	}
	return nr
}
