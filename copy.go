package rex

func (r R) Copy() R {
	s := R{}
	for k, v := range r {
		s[k] = v.Copy()
	}
	return s
}

func (r Relation) Copy() Relation {
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
