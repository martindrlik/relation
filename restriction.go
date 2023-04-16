package rex

func (r *Relation) Restrict(predicate func(tuple map[string]any) bool) *Relation {
	s := NewRelation()
	for _, r := range r.relations {
		for _, t := range r.tuples {
			for _, t := range t {
				if predicate(t) {
					s.InsertTuple(t)
				}
			}
		}
	}
	return s
}
