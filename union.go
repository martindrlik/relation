package rex

func Union(t1, t2 *Table) (*Table, error) {
	if !t1.Schema().IsEqual(t2.Schema()) {
		return nil, ErrSchemaMismatch
	}
	nt := NewTable(t1.SchemaInOrder()...)
	add := func(t T) { nt.Add(t) }
	t1.forEach(add)
	t2.forEach(add)
	return nt, nil
}
