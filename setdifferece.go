package rex

func (r *Relation) SetDifference(s *Relation) *Relation {
	t := NewRelation()
	for k, r := range r.relations {
		if s, ok := s.relations[k]; ok {
			for _, v := range r.setDifference(s) {
				t.InsertTuple(v)
			}
		}
	}
	return t
}

func (r *relation) setDifference(s *relation) []tuple {
	t := []tuple{}
	for _, r := range r.tuples {
		if !s.hasTuple(r) {
			t = append(t, r)
		}
	}
	return t
}
