package main

type Column struct {
	Name string
	Data []any
}

func (column *Column) Append(value any) {
	column.Data = append(column.Data, value)
}

func (column *Column) Set(value any) {
	column.Data[len(column.Data)-1] = value
}
