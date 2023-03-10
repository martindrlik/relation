package rex

func (table *Table) Project(columns ...string) *Table {
	p := Table{}
	m := map[string]struct{}{}
	for _, col := range columns {
		m[col] = struct{}{}
	}
	for _, col := range table.columns {
		if _, ok := m[col.Name]; ok {
			p.columns = append(p.columns, col)
		}
	}
	return &p
}
