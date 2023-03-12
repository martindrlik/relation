package rex

func (t *Table) Project(columns ...string) *Table {
	tm := t.mapColumnByName()
	pc := make([]Column, 0, len(columns))
	for _, cn := range columns {
		if co, ok := tm[cn]; ok {
			pc = append(pc, co)
		}
	}
	return &Table{columns: pc}
}
