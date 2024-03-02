package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/martindrlik/rex/table"
)

func (rex *rex) loadFile(name string) {
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rex.loadReader(path.Base(name), f)
}

func (rex *rex) loadReader(name string, r io.Reader) {
	dec := json.NewDecoder(r)
	for {
		m, err := tryDecode[map[string]any](dec)
		if err != nil && err != io.EOF {
			fmt.Printf("failed to decode: %v\n", err)
		}

		rex.loadTable(name, m)

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

func (rex *rex) loadTable(name string, m map[string]any) {
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
	x, err := table.NewTable(fmap(s, func(a any) string { return a.(string) })...)
	if err != nil {
		fmt.Printf("failed to load table %s: %v\n", name, err)
		return
	}
	for _, row := range rows {
		err := x.Append(row.(map[string]any))
		if err != nil {
			fmt.Printf("failed to append tuple to table %s: %v\n", name, err)
			return
		}
	}
	rex.ts[name] = x
}

func fmap[T any](s []any, f func(any) T) []T {
	t := make([]T, len(s))
	for i, v := range s {
		t[i] = f(v)
	}
	return t
}
