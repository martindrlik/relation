package rex

// Rename returns new relation with attributes renamed.
func (r *Relation) Rename(m map[string]string) *Relation {
	nr := NewRelation()
	for _, s := range r.relations {
		for _, u := range s {
			nr.Insert(u.Rename(m))
		}
	}
	return nr
}

func (t Tuple) Rename(m map[string]string) Tuple {
	r := map[string]any{}
	for k, v := range t {
		if nk, ok := m[k]; ok {
			r[nk] = v
		} else {
			r[k] = v
		}
	}
	return r
}
