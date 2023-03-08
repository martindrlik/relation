package store

type Column struct {
	Name string
	data []any
}

func (column *Column) At(i int) (v any, ok bool) {
	if l := len(column.data); i >= 0 && i < l {
		return column.data[i], true
	}
	return
}

func (column *Column) Len() int {
	return len(column.data)
}

func (column *Column) Insert(value any) {
	column.data = append(column.data, value)
}

func (column *Column) Delete(i int) {
	last := column.Len() - 1
	if i >= 0 && i <= last {
		if i < last {
			column.data[i] = column.data[last]
		}
		column.data = column.data[:last]
	}
}
