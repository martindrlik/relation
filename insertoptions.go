package rex

import (
	"encoding/json"
	"strings"
)

type InsertOptions struct {
	src map[string]any
}

func buildInsertOptions(options ...func(*InsertOptions) error) (*InsertOptions, error) {
	o := &InsertOptions{}
	for _, option := range options {
		err := option(o)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

func String(s string) func(*InsertOptions) error {
	return func(i *InsertOptions) error {
		i.src = map[string]any{}
		dec := json.NewDecoder(strings.NewReader(s))
		return dec.Decode(&i.src)
	}
}

func Map(m map[string]any) func(*InsertOptions) error {
	return func(i *InsertOptions) error {
		i.src = map[string]any{}
		for k, v := range m {
			i.src[k] = v
		}
		return nil
	}
}
