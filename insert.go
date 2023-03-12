package rex

import (
	"encoding/json"
	"sort"
)

func (t *Table) InsertOne(s string) *Table {
	src := map[string]any{}
	err := json.Unmarshal([]byte(s), &src)
	if err != nil {
		panic(err)
	}
	ri := t.dataLen()
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
