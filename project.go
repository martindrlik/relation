package store

func (table *Table) Project(columns ...string) *Table {
	p := Table{}
	m := map[string]struct{}{}
	for _, column := range columns {
		m[column] = struct{}{}
	}
	for _, column := range table.columns {
		if _, ok := m[column.Name]; ok {
			p.columns = append(p.columns, column)
		}
	}
	return &p
}
