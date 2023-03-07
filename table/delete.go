package table

import "github.com/martindrlik/rex/schema"

func (t *Table) Delete(u map[string]any) {
	us := schema.FromTuple(u)
	r, ok := t.Relation(us)
	if ok {
		r.Delete(u)
	}
}
