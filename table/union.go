package table

func (t *Table) Union(u *Table) *Table {
	x := New()
	for _, tuple := range t.tuples {
		x.Add(tuple)
	}
	for _, tuple := range u.tuples {
		x.Add(tuple)
	}
	return x
}
