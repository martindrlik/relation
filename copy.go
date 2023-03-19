package rex

func (r R) copy() R {
	s := R{}
	for k, v := range r {
		s[k] = v.copy()
	}
	return s
}

func (r Relation) copy() Relation {
	s := Relation{}
	s.attributes = make([]string, len(r.attributes))
	s.tuples = make(Tuples, len(r.tuples))
	copy(s.attributes, r.attributes)
	for i, t := range r.tuples {
		s.tuples[i] = make(Tuple, len(t))
		copy(s.tuples[i], t)
	}
	return s
}
