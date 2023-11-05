package rex

type Relation struct {
	attrs
	relations
}

type (
	relations map[attrk]relation
	relation  []Tuple
)

// NewRelation returns new empty relation.
func NewRelation() *Relation {
	return &Relation{relations: relations{}}
}

// Contains returns true if relation contains tuple given by map m.
func (r *Relation) Contains(t map[string]any) bool {
	k := Tuple(t).attrs().key()
	return r.relations[k].contains(t)
}

func (r relation) contains(t map[string]any) bool {
	for _, u := range r {
		if u.equal(t) {
			return true
		}
	}
	return false
}

func (r *Relation) Insert(ts ...map[string]any) *Relation {
	for _, t := range ts {
		attrs := Tuple(t).attrs()
		k := attrs.key()
		if s, ok := r.relations[k]; !ok {
			r.attrs = r.attrs.join(attrs)
			r.relations[k] = []Tuple{t}
		} else {
			r.relations[k] = s.insert(t)
		}
	}
	return r
}

func (r relation) insert(t Tuple) relation {
	if !r.contains(t) {
		return append(r, t)
	}
	return r
}

func (r *Relation) Each(f func(map[string]any) error) error {
	for _, s := range r.relations {
		for _, t := range s {
			err := f(t)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
