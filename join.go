package rex

import "reflect"

func (left *Table) NaturalJoin(right *Table) *Table {
	lci, rci := left.colIntersect(right)
	lcsd := left.colSetDiff(&Table{columns: lci})
	rcsd := right.colSetDiff(&Table{columns: rci})
	src := append(lci, append(lcsd, rcsd...)...)
	ri := rowIntersect(lci, rci)
	t := Table{columns: src}
	for i := range t.columns {
		t.columns[i].data = make([]any, 0, len(ri))
	}
	for _, i := range ri {
		for j, col := range src {
			t.columns[j].data = append(t.columns[j].data, col.data[i])
		}
	}
	return &t
}

func (left *Table) colIntersect(right *Table) (lci, rci []Column) {
	lm := left.colSet()
	for _, rc := range right.columns {
		if lc, ok := lm[rc.Name]; ok {
			lci = append(lci, lc)
			rci = append(rci, rc)
		}
	}
	return
}

func rowIntersect(left, right []Column) (ri []int) {
	for i, lv := range left {
		if reflect.DeepEqual(lv, right[i]) {
			ri = append(ri, i)
		}
	}
	return
}

func (left *Table) colSetDiff(right *Table) (csd []Column) {
	rm := right.colSet()
	for _, lc := range left.columns {
		if _, ok := rm[lc.Name]; !ok {
			csd = append(csd, lc)
		}
	}
	return
}
