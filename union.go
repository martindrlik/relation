package rex

func (r R) Union(s R) R {
	r = r.copy()
	sextra := map[string]Relation{}
	for k, v := range s {
		if rv, ok := r[k]; ok {
			for _, t := range v.tuples {
				rv.tuples.insert(t)
			}
			r[k] = rv
		} else {
			sextra[k] = v
		}
	}
	for k, v := range sextra {
		r[k] = v.copy()
	}
	return r
}
