package rex

import (
	"encoding/json"
	"sort"
)

type Table struct {
	columns columns
}

func (t *Table) DataLen() int {
	if len(t.columns) > 0 {
		return t.columns[0].dataLen()
	}
	return 0
}

func (t *Table) InsertOne(s string) *Table {
	src := map[string]any{}
	err := json.Unmarshal([]byte(s), &src)
	if err != nil {
		panic(err)
	}
	ri := t.DataLen()
	// fill existing columns
	for i, co := range t.columns {
		if v, ok := src[co.name]; ok {
			t.columns[i].insertData(v)
			delete(src, co.name)
		} else {
			t.columns[i].insertData(Empty{})
		}
	}
	// add new columns
	for fn, fv := range src {
		data := make([]any, ri, ri+1)
		for i := 0; i < ri; i++ {
			data[i] = Empty{}
		}
		data = append(data, fv)
		t.columns = append(t.columns, column{
			name: fn,
			data: data})
	}
	if len(src) > 0 {
		sort.Sort(t.columns)
	}
	return t
}

func (t *Table) RemoveAt(i int) *Table {
	for j := range t.columns {
		t.columns[j].removeDataAt(i)
	}
	return t
}

func (t *Table) project(pc ...string) []column {
	if len(pc) == 0 {
		return t.columns
	}
	s := make([]column, len(pc))
	for i, n := range pc {
		s[i] = t.columnByName(n)
	}
	return s
}

func (t *Table) columnByName(s string) column {
	for _, c := range t.columns {
		if c.name == s {
			return c
		}
	}
	return column{name: s}
}

func (t *Table) mapColumnByName() map[string]column {
	m := map[string]column{}
	for _, co := range t.columns {
		m[co.name] = co
	}
	return m
}
