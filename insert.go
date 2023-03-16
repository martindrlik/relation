package rex

import (
	"encoding/json"
	"strings"
)

type Insert struct {
	src map[string]any
}

func String(s string) func(*Insert) error {
	return func(i *Insert) error {
		i.src = map[string]any{}
		dec := json.NewDecoder(strings.NewReader(s))
		return dec.Decode(&i.src)
	}
}

func (r R) InsertOne(options ...func(*Insert) error) (R, error) {
	i := &Insert{}
	for _, option := range options {
		err := option(i)
		if err != nil {
			return R{}, err
		}
	}
	t := newRelation(i.src)
	r.insertRelation(t)
	return r, nil
}

func (r R) insertRelation(s Relation) {
	if len(s.tuples) == 0 {
		return
	}
	rk := s.key()
	gr, ok := r[rk]
	if !ok {
		r[rk] = Relation{s.attributes, [][]any{s.tuples[0]}}
		s.tuples = s.tuples[1:]
		gr = r[rk]
	}
	for _, t := range s.tuples {
		gr.insert(t)
	}
	r[rk] = gr
}

func (r *Relation) insert(t []any) {
	if !r.contains(t) {
		r.tuples = append(r.tuples, t)
	}
}
