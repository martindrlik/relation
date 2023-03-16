package rex

func (r R) Union(s R) R {
	r = r.Copy()
	mr := map[string]Relation{}
	for k, sr := range s {
		rr, ok := r[k]
		if !ok {
			mr[k] = sr
			continue
		}
		for _, st := range sr.tuples {
			rr.insertTuple(st)
		}
		r[k] = rr
	}
	for k, v := range mr {
		r[k] = v
	}
	return r
}
