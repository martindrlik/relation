package rex

// Rename returns new relation with attributes renamed.
func (r *Relation) Rename(rename map[string]string) *Relation {
	nr := NewRelation()
	for ak, ts := range *r {
		as := ak.split()
		nas := make(attrs, len(as))
		for i, a := range as {
			if na, ok := rename[a]; ok {
				nas[i] = na
			} else {
				nas[i] = a
			}
		}
		delete(*nr, ak)
		(*nr)[nas.key()] = ts
	}
	return nr
}
