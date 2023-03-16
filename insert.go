package rex

import (
	"encoding/json"
	"sort"
	"strings"
)

type InsertOptions struct {
	src map[string]any
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

func (r R) Insert(options ...func(*InsertOptions) error) (R, error) {
	in := &InsertOptions{}
	for _, option := range options {
		err := option(in)
		if err != nil {
			return R{}, err
		}
	}

	s := Relation{
		attributes: make([]string, 0, len(in.src)),
	}
	for k := range in.src {
		s.attributes = append(s.attributes, k)
	}
	sort.Strings(s.attributes)
	tuple := make(Tuple, 0, len(s.attributes))
	for _, a := range s.attributes {
		e := in.src[a]
		switch x := e.(type) {
		case map[string]any:
			sr := R{}
			_, err := sr.Insert(Map(x))
			if err != nil {
				return r, err
			}
			tuple = append(tuple, sr)
		default:
			tuple = append(tuple, x)
		}
	}
	s.tuples = append(s.tuples, tuple)
	r.insertRelation(s)
	return r, nil
}

func (r R) insertRelation(s Relation) {
	if len(s.tuples) == 0 {
		return
	}
	rk := s.key()
	gr, ok := r[rk]
	if !ok {
		r[rk] = Relation{attributes: s.attributes}
		gr = r[rk]
	}
	for _, t := range s.tuples {
		gr.tuples.insert(t)
	}
	r[rk] = gr
}
