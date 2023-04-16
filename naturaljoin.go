package rex

func (r *Relation) NaturalJoin(s *Relation) *Relation {
	t := NewRelation()
	for _, r := range r.relations {
		for _, s := range s.relations {
			for _, v := range r.naturalJoin(s, sameAttrs(r, s)) {
				t.InsertTuple(v)
			}
		}
	}
	return t
}

func (r *relation) naturalJoin(s *relation, reduce map[string]struct{}) []tuple {
	t := []tuple{}
	for _, r := range r.tuples {
		for _, r := range r {
			rr := reduceTuple(r, reduce)
			for _, s := range s.tuples {
				for _, s := range s {
					if tupleEquals(rr, reduceTuple(s, reduce)) {
						t = append(t, join(r, s))
					}
				}
			}
		}
	}
	return t
}

func sameAttrs(r, s *relation) map[string]struct{} {
	m := map[string]struct{}{}
	rt := r.tuples.first()
	st := s.tuples.first()
	for k := range rt {
		if _, ok := st[k]; ok {
			m[k] = struct{}{}
		}
	}
	return m
}

func reduceTuple(t tuple, reduce map[string]struct{}) tuple {
	v := tuple{}
	for k := range reduce {
		if w, ok := t[k]; ok {
			v[k] = w
		}
	}
	return v
}

func join(a, b tuple) tuple {
	t := tuple{}
	for k, v := range a {
		t[k] = v
	}
	for k, v := range b {
		t[k] = v
	}
	return t
}
