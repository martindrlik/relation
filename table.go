package rex

type Table struct {
	columns columns
}

func (t *Table) project(pc ...string) []column {
	if len(pc) == 0 {
		return t.columns
	}
	s := make([]column, len(pc))
	for i, n := range pc {
		s[i] = t.columnByName(n)
	}
	return s
}

func (t *Table) columnByName(s string) column {
	for _, c := range t.columns {
		if c.name == s {
			return c
		}
	}
	return column{name: s}
}

func (t *Table) mapColumnByName() map[string]column {
	m := map[string]column{}
	for _, co := range t.columns {
		m[co.name] = co
	}
	return m
}

func (t *Table) dataLen() int {
	if t.columns.Len() > 0 {
		return t.columns[0].dataLen()
	}
	return 0
}
