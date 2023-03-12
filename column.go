package rex

type Column struct {
	Name string
	Data []any
}

func (co *Column) At(i int) (f Field, ok bool) {
	if l := len(co.Data); i >= 0 && i < l {
		return Field{co.Name, co.Data[i]}, true
	}
	return
}

func (co *Column) Len() int {
	return len(co.Data)
}

func (co *Column) Insert(value any) {
	co.Data = append(co.Data, value)
}

func (co *Column) RemoveAt(i int) {
	last := co.Len() - 1
	if i >= 0 && i <= last {
		if i < last {
			co.Data[i] = co.Data[last]
		}
		co.Data = co.Data[:last]
	}
}
