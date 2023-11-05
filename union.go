package rex

// Union returns new relation with tuples from all relations.
func Union(a ...*Relation) *Relation {
	r := NewRelation()
	for _, a := range a {
		a.Each(func(m map[string]any) error {
			r.Insert(m)
			return nil
		})
	}
	return r
}
