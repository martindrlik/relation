package rex

import "github.com/martindrlik/rex/schema"

func Union(t1, t2 *Table) *Table {
	if !schema.IsEqual(t1.Schema(), t2.Schema()) {
		return &Table{}
	}
	tbl := NewTable(t1.SchemaInOrder()...)
	add := func(t T) { tbl.Add(t) }
	t1.forEach(add)
	t2.forEach(add)
	return tbl
}
