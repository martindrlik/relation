package persist

import (
	"encoding/json"
	"io"
	"reflect"

	"github.com/martindrlik/rex/table"
)

func Load(r io.Reader) (*table.Table, error) {
	dec := json.NewDecoder(r)
	tt := []map[string]any{}
	if err := dec.Decode(&tt); err != nil {
		return nil, err
	}
	schema := map[string]reflect.Type{}
	for _, t := range tt {
		for k, v := range t {
			schema[k] = reflect.TypeOf(v)
		}
		break
	}
	t, err := table.New(schema)
	if err != nil {
		return nil, err
	}
	for _, u := range tt {
		if err := t.Add(u); err != nil {
			return nil, err
		}
	}
	return t, nil
}
