package rex

import (
	"encoding/json"
)

type Table struct {
	columns []Column
}

func (t *Table) At(i int) (u []any) {
	if i < 0 || i >= t.Len() {
		return
	}
	for _, col := range t.columns {
		v, _ := col.At(i)
		u = append(u, v)
	}
	return
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
	for i, col := range t.columns {
		if v, ok := src[col.Name]; ok {
			t.columns[i].Insert(v)
			delete(src, col.Name)
		} else {
			t.columns[i].Insert(Empty{})
		}
	}
	// add new columns
	for name, v := range src {
		data := make([]any, insertingRowIndex, insertingRowIndex+1)
		for i := 0; i < insertingRowIndex; i++ {
			data[i] = Empty{}
		}
		data = append(data, v)
		t.columns = append(t.columns, Column{
			Name: name,
			data: data})
	}
	return t
}

func (t *Table) RemoveAt(i int) *Table {
	for j := range t.columns {
		t.columns[j].RemoveAt(i)
	}
	return t
}

func (t *Table) colSet() map[string]Column {
	m := map[string]Column{}
	for _, col := range t.columns {
		m[col.Name] = col
	}
	return m
}
