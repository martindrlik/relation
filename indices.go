package rex

import (
	"reflect"
	"sort"
)

func (t *Table) indices(s *Select) []int {
	if len(s.whereMap) == 0 {
		return t.allIndices()
	}
	m := make(map[int]struct{})
	for _, c := range t.columns {
		if wv, ok := s.whereMap[c.name]; ok {
			for i, cv := range c.data {
				if reflect.DeepEqual(wv, cv) {
					m[i] = struct{}{}
				}
			}
		}
	}
	result := make([]int, 0, len(m))
	for i := range m {
		result = append(result, i)
	}
	sort.Sort(sort.IntSlice(result))
	return result
}

func (t *Table) allIndices() []int {
	if len(t.columns) == 0 {
		return nil
	}
	ri := make([]int, t.columns[0].dataLen())
	for i := range ri {
		ri[i] = i
	}
	return ri
}
