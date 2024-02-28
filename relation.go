package rex

type R []T

func (r *R) Schema() Schema {
	for _, t := range *r {
		return newSchemaTuple(t)
	}
	return newSchema()
}

func (r *R) IsEmpty() bool { return len(*r) == 0 }

func (r *R) Equal(other *R) bool {
	if len(*r) != len(*other) {
		return false
	}
	for i, t := range *r {
		if !t.Equal((*other)[i]) {
			return false
		}
	}
	return true
}

func (r *R) Add(t T) (*R, error) {
	if !r.IsEmpty() && !t.Schema().IsEqual((*r)[0].Schema()) {
		return nil, ErrSchemaMismatch
	}
	if !r.Contain(t) {
		*r = append(*r, t)
	}
	return r, nil
}

func (r *R) Contain(t T) bool {
	for _, v := range *r {
		if v.Equal(t) {
			return true
		}
	}
	return false
}
