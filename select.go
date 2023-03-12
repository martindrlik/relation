package rex

import "encoding/json"

type Select struct {
	whereMap map[string]interface{}
	pcols    []string
}

func Where(restrict string) func(*Select) {
	wm := map[string]interface{}{}
	err := json.Unmarshal([]byte(restrict), &wm)
	if err != nil {
		panic(err)
	}
	return func(s *Select) {
		s.whereMap = wm
	}
}

func Project(columns ...string) func(*Select) {
	return func(s *Select) {
		s.pcols = columns
	}
}

func (t *Table) Select(options ...func(*Select)) [][]any {
	s := &Select{}
	for _, option := range options {
		option(s)
	}
	si := t.indices(s)
	pc := t.project(s.pcols)
	rs := make([][]any, 0, len(si))
	for _, ri := range si {
		row := make([]any, len(pc))
		for ci, data := range pc {
			row[ci] = data[ri]
		}
		rs = append(rs, row)
	}
	return rs
}
