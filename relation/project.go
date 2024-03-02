package relation

import (
	"github.com/martindrlik/rex/schema"
)

func (r *Relation) Project(attributes ...string) (*Relation, error) {
	if len(attributes) == 0 {
		return nil, ErrMissingSchema
	}
	s := schema.New(attributes...)
	x, err := New(r.t[0].Project(s))
	if err != nil {
		return nil, err
	}
	for _, t := range r.t[1:] {
		x.append(t.Project(s))
	}
	return x, nil
}
