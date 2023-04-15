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

func (r *relation) naturalJoin(s *relation, reduce map[string]struct{}) []map[string]any {
	t := []map[string]any{}
	for _, r := range r.tuples {
		rr := reduceTuple(r, reduce)
		for _, s := range s.tuples {
			if tupleEquals(rr, reduceTuple(s, reduce)) {
				t = append(t, join(r, s))
			}
		}
	}
	return t
}

func sameAttrs(r, s *relation) map[string]struct{} {
	m := map[string]struct{}{}
	for k := range r.tuples[0] {
		if _, ok := s.tuples[0][k]; ok {
			m[k] = struct{}{}
		}
	}
	return m
}

func reduceTuple(tuple map[string]any, reduce map[string]struct{}) map[string]any {
	m := map[string]any{}
	for k := range reduce {
		if v, ok := tuple[k]; ok {
			m[k] = v
		}
	}
	return m
}

func join(a, b map[string]any) map[string]any {
	m := map[string]any{}
	for k, v := range a {
		m[k] = v
	}
	for k, v := range b {
		m[k] = v
	}
	return m
}
