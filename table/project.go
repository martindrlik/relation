package table

func (t *Table) Project(schema ...string) *Table {
	x := New()
	add := func() func(map[string]any) {
		ptf := projectTuple(schema...)
		return func(tuple map[string]any) {
			tuple, ok := ptf(tuple)
			if ok {
				x.Add(tuple)
			}
		}
	}()
	for _, tuple := range t.tuples {
		add(tuple)
	}
	return x
}

func projectTuple(schema ...string) func(map[string]any) (map[string]any, bool) {
	return func(tuple map[string]any) (map[string]any, bool) {
		x := make(map[string]any)
		for _, k := range schema {
			v, ok := tuple[k]
			if !ok {
				return nil, false
			}
			x[k] = v
		}
		return x, true
	}
}
