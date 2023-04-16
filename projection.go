package rex

func (r *Relation) Project(attributes ...string) *Relation {
	s := NewRelation()
	for _, r := range r.relations {
		for _, t := range r.tuples {
			for _, t := range t {
				m := tuple{}
				for _, a := range attributes {
					if v, ok := t[a]; ok {
						m[a] = v
					}
				}
				s.InsertTuple(m)
			}
		}
	}
	return s
}
