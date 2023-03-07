package main

import "encoding/json"

type Table struct {
	columns []Column
}

func (table *Table) firstColumn() (*Column, bool) {
	if len(table.columns) > 0 {
		return &table.columns[0], true
	}
	return nil, false
}

func (table *Table) ColumnByName(name string) (*Column, bool) {
	for i, column := range table.columns {
		if name == column.Name {
			return &table.columns[i], true
		}
	}
	return nil, false
}

func (table *Table) appendColumn(name string) *Column {
	var data []any
	if column, ok := table.firstColumn(); ok {
		for _ = range column.Data {
			data = append(data, NoValue{})
		}
	}
	table.columns = append(table.columns, Column{Name: name})
	return &table.columns[len(table.columns)-1]
}

func (table *Table) Append(s string) error {
	m := map[string]any{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return err
	}
	for i, column := range table.columns {
		if value, ok := m[column.Name]; ok {
			table.columns[i].Append(value)
			delete(m, column.Name)
		} else {
			table.columns[i].Append(NoValue{})
		}
	}
	for name, value := range m {
		table.appendColumn(name).Append(value)
	}
	return nil
}

func (table *Table) Row(n int) ([]any, error) {
	column, ok := table.firstColumn()
	if !ok || n < 0 || n >= len(column.Data) {
		return nil, NoData
	}
	data := make([]any, len(table.columns))
	for i, column := range table.columns {
		data[i] = column.Data[n]
	}
	return data, nil
}
