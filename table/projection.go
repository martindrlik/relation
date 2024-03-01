package table

func (t *Table) Project(attributes ...string) *Table {
	nt := NewTable(attributes...)
	for _, r := range t.r {
		nr := r.Project(nt.Schema())
		for _, t := range nr.Tuples() {
			nt.Append(t)
		}
	}
	return nt
}
