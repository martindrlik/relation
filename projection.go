package rex

func (r *Relation) Project(attributes ...string) *Relation {
	s := NewRelation()
	for _, r := range r.relations {
		for _, t := range r.tuples {
			m := map[string]any{}
			for _, a := range attributes {
				if v, ok := t[a]; ok {
					m[a] = v
				}
			}
			s.InsertTuple(m)
		}
	}
	return s
}
