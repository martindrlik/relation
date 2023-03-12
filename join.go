package rex

func NaturalJoin(tables ...*Table) *Table {
	ci := columnsIntersect(tables...)
	if len(ci) <= 0 {
		return nil
	}
	srcm := columns{}
	add := func(c columns) {
		for name, data := range c {
			srcm[name] = data
		}
	}
	add(ci[0])
	for i, t := range tables {
		add(t.columnsExcept(ci[i]))
	}
	dstm := columns{}
	for _, ri := range rowIntersect(ci[0], ci[1]) {
		for name, data := range srcm {
			dstm[name] = append(dstm[name], data[ri])
		}
	}
	return &Table{columns: dstm}
}
