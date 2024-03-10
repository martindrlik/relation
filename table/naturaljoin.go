package table

func (t *Table) NaturalJoin(u *Table) *Table {
	x := New()
	for _, tuple := range t.CompleteTuples() {
		for _, other := range u.CompleteTuples() {
			if T(tuple).EqualsOnCommon(other) {
				x.Add(T(tuple).Merge(other))
			}
		}
	}
	return x
}
