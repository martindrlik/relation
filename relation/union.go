package relation

func (r *Relation) Union(others ...*Relation) (*Relation, error) {
	for _, o := range others {
		if !r.Schema().IsEqual(o.Schema()) {
			return nil, ErrSchemaMismatch
		}
	}

	x := must(r.Project(r.Schema().Attributes()...))

	for _, o := range others {
		x.appendRelation(o)
	}

	return x, nil
}
