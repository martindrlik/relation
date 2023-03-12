package rex

import (
	"reflect"
	"sort"
)

func (sel *Select) rowIndices(t *Table) []int {
	if len(sel.whereMap) == 0 {
		ri := make([]int, t.Len())
		for i := range ri {
			ri[i] = i
		}
		return ri
	}
	rm := map[int]struct{}{}
	for _, co := range t.columns {
		if wv, ok := sel.whereMap[co.Name]; ok {
			for j, v := range co.Data {
				if reflect.DeepEqual(wv, v) {
					rm[j] = struct{}{}
				}
			}
		}
	}
	ri := make([]int, 0, len(rm))
	for i := range rm {
		ri = append(ri, i)
	}
	sort.Sort(sort.IntSlice(ri))
	return ri
}
