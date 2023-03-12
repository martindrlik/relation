package rex

import "encoding/json"

type Select struct {
	whereMap map[string]interface{}
	pcolumns []string
}

func Where(w string) func(*Select) {
	wm := map[string]interface{}{}
	err := json.Unmarshal([]byte(w), &wm)
	if err != nil {
		panic(err)
	}
	return func(s *Select) {
		s.whereMap = wm
	}
}

func Project(columns ...string) func(*Select) {
	return func(s *Select) {
		s.pcolumns = columns
	}
}

func (t *Table) Select(options ...func(*Select)) [][]any {
	s := &Select{}
	for _, option := range options {
		option(s)
	}
	ri := t.indices(s)
	pc := t.project(s.pcolumns...)
	rows := make([][]any, 0, len(ri))
	for _, i := range ri {
		row := make([]any, len(pc))
		for j, c := range pc {
			row[j] = c.dataAt(i)
		}
		rows = append(rows, row)
	}
	return rows
}
