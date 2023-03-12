package rex

func (t *Table) Delete(options ...func(*Select)) int {
	s := &Select{}
	for _, option := range options {
		option(s)
	}
	ri := t.indices(s)
	ln := len(ri)
	for i := ln - 1; i >= 0; i-- {
		for j := range t.columns {
			t.columns[j].removeDataAt(ri[i])
		}
	}
	return ln
}
