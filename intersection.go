package rex

func (r *Relation) Intersect(s *Relation) *Relation {
	return r.SetDifference(r.SetDifference(s))
}
