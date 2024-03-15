package load

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/martindrlik/rex/table"
)

func TableFiles(names ...string) ([]*table.Table, error) {
	tables := []*table.Table{}
	for _, name := range names {
		t, err := loadTableFile(name)
		if err != nil {
			return nil, fmt.Errorf("failed to load file %s: %w", name, err)
		}
		tables = append(tables, t)
	}
	return tables, nil
}

func loadTableFile(name string) (*table.Table, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	tuples := []map[string]any{}
	if err := dec.Decode(&tuples); err != nil {
		return nil, err
	}
	return table.New().Add(tuples...), nil
}
