package rex

import "reflect"

func (left *Table) NaturalJoin(right *Table) *Table {
	lci, rci := left.colIntersect(right)
	lcsd := left.colSetDiff(&Table{columns: lci})
	rcsd := right.colSetDiff(&Table{columns: rci})
	src := append(lci, append(lcsd, rcsd...)...)
	dst := make([]column, len(src))
	ri := rowIntersect(lci, rci)
	for i, col := range src {
		dst[i] = column{
			name: col.name,
			data: make([]any, 0, len(ri)),
		}
	}
	result := Table{columns: dst}
	for _, i := range ri {
		for j, c := range src {
			result.columns[j].insertData(c.data[i])
		}
	}
	return &result
}

func (left *Table) colIntersect(right *Table) (lci, rci []column) {
	lm := left.mapColumnByName()
	for _, rc := range right.columns {
		if lc, ok := lm[rc.name]; ok {
			lci = append(lci, lc)
			rci = append(rci, rc)
		}
	}
	return
}

func rowIntersect(left, right []column) (ri []int) {
	for i, lc := range left {
		for j, lv := range lc.data {
			if reflect.DeepEqual(lv, right[i].data[j]) {
				ri = append(ri, j)
			}
		}
	}
	return
}

func (left *Table) colSetDiff(right *Table) (csd []column) {
	rm := right.mapColumnByName()
	for _, lc := range left.columns {
		if _, ok := rm[lc.name]; !ok {
			csd = append(csd, lc)
		}
	}
	return
}
