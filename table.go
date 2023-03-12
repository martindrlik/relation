package rex

import "sort"

type Table struct {
	columns columns
}

func (t *Table) project(pc []string) [][]any {
	pc = func() []string {
		if len(pc) > 0 {
			return pc
		}
		o := make([]string, 0, len(t.columns))
		for name := range t.columns {
			o = append(o, name)
		}
		sort.Strings(o)
		return o
	}()
	s := make([][]any, len(pc))
	for i, name := range pc {
		s[i] = t.columns[name]
	}
	return s
}

func (t *Table) dataLen() int {
	for _, data := range t.columns {
		return len(data)
	}
	return 0
}
