package rex

import (
	"encoding/json"
	"io"
)

func (r *Relation) InsertOneJson(s io.Reader) *Relation {
	dec := json.NewDecoder(s)
	m := map[string]any{}
	err := dec.Decode(&m)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return r.InsertTuple(m)
}

func (r *Relation) InsertManyJson(s io.Reader) *Relation {
	dec := json.NewDecoder(s)
	m := []map[string]any{}
	err := dec.Decode(&m)
	if err != nil && err != io.EOF {
		panic(err)
	}
	for _, m := range m {
		r.InsertTuple(m)
	}
	return r
}
