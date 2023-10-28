package rex

// Difference returns new relation with tuples from left relation not present in right relation.
func Difference(left, right *Relation) *Relation {
	r := NewRelation()
	for k, lr := range *left {
		rr, ok := (*right)[k]
		if !ok {
			cr := make(relation, len(lr))
			copy(cr, lr)
			(*r)[k] = cr
			continue
		}
		for _, lt := range lr {
			if !rr.contains(lt) {
				r.insertTuplex(k, lt)
			}
		}
	}
	return r
}
