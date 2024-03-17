package persist

import (
	"encoding/json"
	"io"

	"github.com/martindrlik/rex/table"
)

func Load(r io.Reader) (*table.Table, error) {
	dec := json.NewDecoder(r)
	tuples := []map[string]any{}
	if err := dec.Decode(&tuples); err != nil {
		return nil, err
	}
	return table.New().Add(tuples...), nil
}
