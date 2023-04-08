package rex

type WhereOptions struct {
	greaterThanName  string
	greaterThanValue any
}

func (r R) Where(options ...func(*WhereOptions)) R {
	wo := new(WhereOptions)
	for _, o := range options {
		o(wo)
	}
	s := R{}
	for _, r := range r {
		ai := r.attri()
		i, ok := ai[wo.greaterThanName]
		if !ok {
			continue
		}
		u := Relation{}
		for _, t := range r.tuples {
			if !greaterThan(t[i], wo.greaterThanValue) {
				continue
			}
			if len(u.attributes) == 0 {
				u.attributes = r.attributes
			}
			u.tuples = append(u.tuples, t)
		}
		if len(u.tuples) > 0 {
			s.insertRelation(u)
		}
	}
	return s
}

func GreaterThan[V comparable](name string, value V) func(*WhereOptions) {
	return func(wo *WhereOptions) {
		wo.greaterThanName = name
		wo.greaterThanValue = value
	}
}

func greaterThan(a, b any) bool {
	ax, ok := a.(float64)
	if !ok {
		return false
	}
	bx, ok := b.(float64)
	if !ok {
		return false
	}
	return ax > bx
}
