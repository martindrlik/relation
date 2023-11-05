package rex

func (r *Relation) Projection(a ...string) *Relation {
	nr := NewRelation()
	for _, s := range r.relations {
		for _, u := range s {
			nr.Insert(u.Projection(a...))
		}
	}
	return nr
}

func (t Tuple) Projection(a ...string) Tuple {
	r := map[string]any{}
	for _, a := range a {
		if v, ok := t[a]; ok {
			r[a] = v
		}
	}
	return r
}
