package table

import "github.com/martindrlik/rex/schema"

// Rename1 returns a new table with the attribute old renamed to new.
func (t *Table) Rename1(old, new string) (*Table, error) {
	if !t.Schema.Has(old) || t.Schema.Has(new) {
		return nil, schema.ErrMismatch
	}

	s := schema.Schema{}
	for k, v := range t.Schema {
		s[k] = v
	}
	s[new] = s[old]
	delete(s, old)

	w, _ := New(s)

	for _, tr := range t.RelationSet {
		if tr.Schema.Has(old) {
			wr, _ := tr.Rename1(old, new)
			w.RelationSet.Add(wr)
		}
	}

	return w, nil
}
