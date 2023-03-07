package table

import "github.com/martindrlik/rex/schema"

// Delete deletes tuples from the table.
func (t *Table) Delete(u map[string]any) {
	us := schema.FromTuple(u)
	r, ok := t.Relation(us)
	if ok {
		r.Delete(u)
	}
}
