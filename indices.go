package rex

import (
	"reflect"
	"sort"
)

func (t *Table) indices(s *Select) []int {
	if t.columns.Len() == 0 {
		return nil
	}
	if len(s.whereMap) == 0 {
		result := make([]int, t.columns[0].dataLen())
		for i, ln := 0, t.columns[0].dataLen(); i < ln; i++ {
			result[i] = i
		}
		return result
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
	sort.Ints(result)
	return result
}
