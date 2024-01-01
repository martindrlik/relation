package rex

// RS represents a set of relations.
type RS []*R

// New returns a new relation set.
func New() *RS { return &RS{} }

func (rs *RS) Project(as ...string) *RS  { return &RS{rs.R().Project(as...)} }
func (rs *RS) Union(other *RS) *RS       { return &RS{rs.R().Union(other.R())} }
func (rs *RS) Difference(other *RS) *RS  { return &RS{rs.R().Difference(other.R())} }
func (rs *RS) NaturalJoin(other *RS) *RS { return &RS{rs.R().NaturalJoin(other.R())} }

func (rs *RS) Add(t T) {
	r, ok := rs.byTuple(t)
	if !ok {
		r = &R{}
		rs.addRel(r)
	}
	r.Add(t)
}

func (rs *RS) addRel(r *R) {
	*rs = append(*rs, r)
}

func (rs *RS) Remove(t T) {
	if r, ok := rs.byTuple(t); ok {
		r.Remove(t)
		if r.IsEmpty() {
			rs.removeRel(r)
		}
	}
}

func (rs *RS) removeRel(r *R) {
	nrs := make(RS, 0, len(*rs)-1)
	for _, r2 := range *rs {
		if r != r2 {
			nrs.addRel(r2)
		}
	}
	*rs = nrs
}

func (rs *RS) byTuple(t T) (*R, bool) {
	for _, r := range *rs {
		if r.first().IsCompatible(t) {
			return r, true
		}
	}
	return nil, false
}

// R returns a relation that has all the attributes of the relations in the set.
func (rs *RS) R() *R {
	as := rs.attributes()
	for _, r := range *rs {
		if len(r.first()) == len(as) && r.HasAttributes(as...) {
			return r
		}
	}
	return &R{}
}

func (rs *RS) attributes() []string {
	am := map[string]struct{}{}
	for _, r := range *rs {
		for a := range r.first() {
			am[a] = struct{}{}
		}
	}
	as := make([]string, 0, len(am))
	for a := range am {
		as = append(as, a)
	}
	return as
}
