package rex

func (r R) Copy() R {
	c := R{}
	for k, v := range r {
		c[k] = v.Copy()
	}
	return c
}

func (r Relation) Copy() Relation {
	ca := make([]string, len(r.attributes))
	copy(ca, r.attributes)
	ct := make([][]any, len(r.tuples))
	for i, t := range r.tuples {
		ct[i] = make([]any, len(t))
		copy(ct[i], t)
	}
	return Relation{ca, ct}
}
