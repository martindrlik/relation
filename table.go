package rex

import (
	"encoding/json"
)

type Table struct {
	columns []Column
}

func (t *Table) Select(options ...func(*Select)) [][]any {
	sel := &Select{}
	for _, option := range options {
		option(sel)
	}
	ri := sel.rowIndices(t)
	result := make([][]any, len(ri))
	for i, ii := range ri {
		result[i] = make([]any, len(t.columns))
		for j, co := range t.columns {
			result[i][j] = co.Data[ii]
		}
	}
	return result
}

func (t *Table) SelectRange(from, to int, columns ...string) [][]any {
	ln := t.Len()
	if ln == 0 {
		return nil
	}
	if from < 0 || from >= ln {
		from = 0
	}
	if to < 0 || to >= ln {
		to = ln - 1
	}
	tm := t.mapColumnByName()
	fs := make([][]any, 1+to-from)
	for ri := from; ri <= to; ri++ {
		fs[ri] = make([]any, len(columns))
		for ci, cn := range columns {
			if co, ok := tm[cn]; ok {
				f, _ := co.At(ri)
				fs[ri][ci] = f.Value
			}
		}
	}
	return fs
}

func (t *Table) Len() int {
	if len(t.columns) > 0 {
		return t.columns[0].Len()
	}
	return 0
}

func (t *Table) InsertOne(s string) *Table {
	src := map[string]any{}
	err := json.Unmarshal([]byte(s), &src)
	if err != nil {
		panic(err)
	}
	insertingRowIndex := t.Len()
	// fill existing columns
	for i, co := range t.columns {
		if v, ok := src[co.Name]; ok {
			t.columns[i].Insert(v)
			delete(src, co.Name)
		} else {
			t.columns[i].Insert(Empty{})
		}
	}
	// add new columns
	for fn, fv := range src {
		data := make([]any, insertingRowIndex, insertingRowIndex+1)
		for i := 0; i < insertingRowIndex; i++ {
			data[i] = Empty{}
		}
		data = append(data, fv)
		t.columns = append(t.columns, Column{
			Name: fn,
			Data: data})
	}
	return t
}

func (t *Table) RemoveAt(i int) *Table {
	for j := range t.columns {
		t.columns[j].RemoveAt(i)
	}
	return t
}

func (t *Table) mapColumnByName() map[string]Column {
	m := map[string]Column{}
	for _, co := range t.columns {
		m[co.Name] = co
	}
	return m
}
