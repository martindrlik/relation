package store

import "encoding/json"

type Table struct {
	columns []Column
}

func (table *Table) Len() int {
	if len(table.columns) > 0 {
		return table.columns[0].Len()
	}
	return 0
}

func (table *Table) InsertOne(s string) error {
	src := map[string]any{}
	err := json.Unmarshal([]byte(s), &src)
	if err != nil {
		return err
	}
	insertingRowIndex := table.Len()
	// fill existing columns
	for i, column := range table.columns {
		if v, ok := src[column.Name]; ok {
			table.columns[i].Insert(v)
			delete(src, column.Name)
		} else {
			table.columns[i].Insert(NoValue{})
		}
	}
	// add new columns
	for name, v := range src {
		data := make([]any, insertingRowIndex, insertingRowIndex+1)
		for i := 0; i < insertingRowIndex; i++ {
			data[i] = NoValue{}
		}
		data = append(data, v)
		table.columns = append(table.columns, Column{
			Name: name,
			data: data})
	}

	return nil
}

func (table *Table) At(i int) (t []any, ok bool) {
	ok = i >= 0 && i < table.Len()
	if !ok {
		return
	}
	for _, column := range table.columns {
		v, _ := column.At(i)
		t = append(t, v)
	}
	return
}
