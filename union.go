package rex

func (r R) Union(s R) R {
	t := r.Copy()
	v := s.Copy()
	for k := range t {
		if vv, ok := v[k]; ok {
			t.insertRelation(vv)
			delete(v, k)
		}
	}
	for k, vv := range v {
		t[k] = vv
	}
	return t
}
