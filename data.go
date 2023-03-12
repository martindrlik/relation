package rex

import (
	"reflect"
	"sort"
)

func (t *Table) indices(s *Select) []int {
	if dl, wml := t.dataLen(), len(s.whereMap); dl == 0 || wml == 0 {
		all := make([]int, dl)
		for i := range all {
			all[i] = i
		}
		return all
	}
	sm := make(map[int]struct{})
	for name, wv := range s.whereMap {
		if data, ok := t.columns[name]; ok {
			for i, cv := range data {
				if reflect.DeepEqual(wv, cv) {
					sm[i] = struct{}{}
				}
			}
		}
	}
	o := make([]int, 0, len(sm))
	for i := range sm {
		o = append(o, i)
	}
	sort.Ints(o)
	return o
}

func rowIntersect(u, v columns) []int {
	sri := make([]int, 0)
	for name, data := range u {
		for ri, uv := range data {
			if reflect.DeepEqual(uv, v[name][ri]) {
				sri = append(sri, ri)
			}
		}
	}
	return sri
}
