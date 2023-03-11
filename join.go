package rex

import "reflect"

func (left *Table) NaturalJoin(right *Table) *Table {
	lci, rci := left.colIntersect(right)
	lcsd := left.colSetDiff(&Table{columns: lci})
	rcsd := right.colSetDiff(&Table{columns: rci})
	src := append(lci, append(lcsd, rcsd...)...)
	dst := make([]Column, len(src))
	ri := rowIntersect(lci, rci)
	for i, col := range src {
		dst[i] = Column{
			Name: col.Name,
			data: make([]any, 0, len(ri)),
		}
	}
	t := Table{columns: dst}
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
	for i, lc := range left {
		for j, lv := range lc.data {
			if reflect.DeepEqual(lv, right[i].data[j]) {
				ri = append(ri, j)
			}
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
