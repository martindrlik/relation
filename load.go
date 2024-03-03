package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/martindrlik/rex/table"
)

func loadFile(name string) (*table.Table, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decode(f)
}

func decode(r io.Reader) (*table.Table, error) {
	dec := json.NewDecoder(r)
	raw := map[string]any{}
	err := dec.Decode(&raw)
	if err != nil && err != io.EOF {
		return nil, err
	}

	schema, err := decodeSchema(raw)
	if err != nil {
		return nil, err
	}

	x, err := table.New(schema...)
	if err != nil {
		return nil, err
	}

	rows, err := decodeRows(raw, schema)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		if err := x.Append(row); err != nil {
			return nil, err
		}
	}

	return x, nil
}

func decodeSchema(raw map[string]any) ([]string, error) {
	anySchema, ok := raw["schema"].([]any)
	if !ok {
		return nil, errors.New("missing schema")
	}
	schema := []string{}
	for _, s := range anySchema {
		schema = append(schema, fmt.Sprintf("%v", s))
	}
	return schema, nil
}

func decodeRows(raw map[string]any, schema []string) ([]map[string]any, error) {
	anyRows, ok := raw["rows"].([]any)
	if !ok {
		return nil, errors.New("missing rows")
	}
	rows := []map[string]any{}
	for _, anyRow := range anyRows {
		switch row := anyRow.(type) {
		case map[string]any:
			rows = append(rows, row)
		case []any:
			if len(schema) != len(row) {
				return nil, fmt.Errorf("tuple length mismatch: %d != %d", len(schema), len(row))
			}
			x := map[string]any{}
			for i, v := range schema {
				x[v] = row[i]
			}
			rows = append(rows, x)
		default:
			return nil, fmt.Errorf("unsupported row type %v: %T", anyRow, row)
		}
	}
	return rows, nil
}
