package table

func (t *Table) NaturalJoin(u *Table) *Table {
	x := New()
	isCompleteT := t.isCompleteTuple()
	isCompleteU := u.isCompleteTuple()
	for _, tuple := range t.tuples {
		if !isCompleteT(tuple) {
			continue
		}
		for _, other := range u.tuples {
			if !isCompleteU(other) {
				continue
			}
			if T(tuple).EqualsOnCommon(other) {
				x.Add(T(tuple).Merge(other))
			}
		}
	}
	return x
}
