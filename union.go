package rex

func (r R) Union(s R) R {
	r = r.Copy()
	rextra := map[string]Relation{}
	for k, v := range s {
		if rv, ok := r[k]; ok {
			for _, t := range v.tuples {
				rv.tuples.insert(t)
			}
			r[k] = rv
		} else {
			rextra[k] = v
		}
	}
	return r
}
