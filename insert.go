package rex

import (
	"io"
)

func (g R) InsertOne(s io.Reader) R {
	m := mustDecode[map[string]any](s)
	r := newRelation(m)
	g.insertRelation(r)
	return g
}

func (g R) insertRelation(r Relation) {
	if len(r.tuples) == 0 {
		return
	}
	rk := r.key()
	gr, ok := g[rk]
	if !ok {
		g[rk] = Relation{r.attributes, [][]any{r.tuples[0]}}
		r.tuples = r.tuples[1:]
		gr = g[rk]
	}
	for _, t := range r.tuples {
		if !gr.contains(t) {
			gr.tuples = append(gr.tuples, t)
		}
	}
	g[rk] = gr
}
