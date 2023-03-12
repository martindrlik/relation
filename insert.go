package rex

import (
	"encoding/json"
)

func (t *Table) InsertOne(s string) *Table {
	srcm := map[string]any{}
	err := json.Unmarshal([]byte(s), &srcm)
	if err != nil {
		panic(err)
	}
	ri := t.dataLen()
	if t.columns == nil {
		t.columns = columns{}
	}
	// fill existing columns
	for name, data := range t.columns {
		if srcv, ok := srcm[name]; ok {
			t.columns[name] = append(data, srcv)
			delete(srcm, name)
		} else {
			t.columns[name] = append(data, Empty{})
		}
	}
	// add new columns
	for srcname, srcv := range srcm {
		data := make([]any, ri+1)
		for i := 0; i < ri; i++ {
			data[i] = Empty{}
		}
		data[ri] = srcv
		t.columns[srcname] = data
	}
	return t
}
