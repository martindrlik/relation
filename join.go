package rex

func NaturalJoin(u, v *Table) *Table {
	ci := columnsIntersect(u, v)
	if len(ci) <= 0 {
		return nil
	}
	usd := u.columnsExcept(ci[0])
	vsd := v.columnsExcept(ci[1])

	srcm := columns{}
	dstm := columns{}
	add := func(c columns) {
		for name, data := range c {
			srcm[name] = data
		}
	}
	add(ci[0])
	add(usd)
	add(vsd)
	for _, ri := range rowIntersect(ci[0], ci[1]) {
		for name, data := range srcm {
			dstm[name] = append(dstm[name], data[ri])
		}
	}
	return &Table{columns: dstm}
}
