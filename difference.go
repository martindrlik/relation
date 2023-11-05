package rex

// Difference returns new relation with tuples from left relation not present in right relation.
func Difference(left, right *Relation) *Relation {
	r := NewRelation()
	for _, ls := range left.relations {
		for _, lt := range ls {
			if !right.Contains(lt) {
				r.Insert(lt)
			}
		}
	}
	return r
}
