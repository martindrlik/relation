package rex

type Column struct {
	Name string
	data []any
}

func (col *Column) At(i int) (v any, ok bool) {
	if l := len(col.data); i >= 0 && i < l {
		return col.data[i], true
	}
	return
}

func (col *Column) Len() int {
	return len(col.data)
}

func (col *Column) Insert(value any) {
	col.data = append(col.data, value)
}

func (col *Column) RemoveAt(i int) {
	last := col.Len() - 1
	if i >= 0 && i <= last {
		if i < last {
			col.data[i] = col.data[last]
		}
		col.data = col.data[:last]
	}
}
