package table

func Union(t, v *Table) (*Table, error) {
	if !t.Schema().IsEqual(v.Schema()) {
		return nil, ErrMustBeUnionCompatible
	}
	nt := NewTable(t1.Schema()...)
	add := func(t T) { nt.Add(t) }
	t1.forEach(add)
	t2.forEach(add)
	return nt, nil
}
