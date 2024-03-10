package table

func (t *Table) SetDifference(u *Table) *Table {
	x := New()
	isComplete := t.isCompleteTuple()
	for _, tuple := range t.tuples {
		if !isComplete(tuple) || !u.Tuples().Contains(tuple) {
			x.Add(tuple)
		}
	}
	return x
}
