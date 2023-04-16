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

func (r *relation) union(s *relation) []tuple {
	t := []tuple{}
	add := func(tuples tuples) {
		for _, v := range tuples {
			for _, v := range v {
				t = append(t, v.shallowCopy())
			}
		}
	}
	add(r.tuples)
	add(s.tuples)
	return t
}
