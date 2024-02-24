package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/martindrlik/rex"
)

func (state *state) loadCmd(a args) {
	switch len(a) {
	case 0:
	default:
		state.loadFile(a.first())
	}
}

func (state *state) loadFile(name string) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	state.loadReader(path.Base(name), f)
}

func (state *state) loadReader(name string, r io.Reader) {
	dec := json.NewDecoder(r)
	for {
		m, err := tryDecode[map[string]any](dec)
		if err != nil && err != io.EOF {
			fmt.Printf("failed to decode: %v\n", err)
		}

		state.loadTable(name, m)

		if err != nil {
			break
		}
	}
}

func tryDecode[T any](dec *json.Decoder) (t T, err error) {
	if err := dec.Decode(&t); err != nil {
		return t, err
	}
	return t, nil
}

func (state *state) loadTable(name string, m map[string]any) {
	if len(m) == 0 {
		return
	}
	s, ok := m["schema"].([]any)
	if !ok {
		fmt.Printf("failed to load table %s: schema not found\n", name)
		return
	}
	rows, ok := m["rows"].([]any)
	if !ok {
		fmt.Printf("failed to load table %s: rows not found\n", name)
		return
	}
	table := rex.NewTable(fmap(s, func(a any) string {
		return a.(string)
	})...)
	for _, row := range rows {
		table.Add(row.(map[string]any))
	}
	state.tables[name] = table
}

func fmap[T any](s []any, f func(any) T) []T {
	t := make([]T, len(s))
	for i, v := range s {
		t[i] = f(v)
	}
	return t
}
