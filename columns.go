package rex

type columns map[string][]any

func (c columns) removeDataAt(i int) {
	for name, data := range c {
		last := len(data) - 1
		if i < last {
			data[i] = data[last]
		}
		c[name] = data[:last]
	}
}

func (t *Table) columnsExcept(except columns) columns {
	cem := columns{}
	for name, data := range t.columns {
		if _, ok := except[name]; !ok {
			cem[name] = data
		}
	}
	return cem
}

func columnsIntersect(tables ...*Table) []columns {
	im := map[string]struct{}{}
	for name := range tables[0].columns {
		ok := false
		for i := 1; i < len(tables); i++ {
			if _, ok = tables[i].columns[name]; !ok {
				break
			}
		}
		if ok {
			im[name] = struct{}{}
		}
	}
	if len(im) <= 0 {
		return nil
	}
	r := make([]columns, len(tables))
	for i, t := range tables {
		r[i] = columns{}
		for name := range im {
			r[i][name] = t.columns[name]
		}
	}
	return r
}
