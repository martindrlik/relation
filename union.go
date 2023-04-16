package rex

func (r *Relation) Union(s *Relation) *Relation {
	t := NewRelation()
	for k, r := range r.relations {
		if s, ok := s.relations[k]; ok {
			for _, v := range r.union(s) {
				t.InsertTuple(v)
			}
		}
	}
	return t
}

func (r *relation) union(s *relation) []map[string]any {
	t := []map[string]any{}
	add := func(tuples []tuple) {
		for _, v := range tuples {
			t = append(t, v.shallowCopy())
		}
	}
	add(r.tuples)
	add(s.tuples)
	return t
}
