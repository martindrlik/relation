package rex

// R represents a relation.
type R []T

func (r *R) Add(t T) {
	if !r.Contains(t) {
		r.add(t)
	}
}

func (r *R) add(t T) {
	*r = append(*r, t)
}

func (r *R) Remove(t T) {
	if i := r.index(t); i >= 0 {
		r.remove(i)
	}
}

func (r *R) remove(i int) {
	nr := make(R, 0, len(*r)-1)
	nr = append(nr, (*r)[:i]...)
	nr = append(nr, (*r)[i+1:]...)
	*r = nr
}

func (r *R) IsEmpty() bool { return r.Len() == 0 }
func (r *R) Len() int      { return len(*r) }

func (r *R) IsEqual(other *R) bool {
	return r.Len() == other.Len() && r.isEqual(other)
}

func (r *R) isEqual(other *R) bool {
	for _, t := range *r {
		if !other.Contains(t) {
			return false
		}
	}
	return true
}

func (r *R) Contains(t T) bool { return r.index(t) >= 0 }

func (r *R) index(t T) int {
	for i, t2 := range *r {
		if t.IsEqual(t2) {
			return i
		}
	}
	return -1
}

func (r *R) HasAttributes(as ...string) bool {
	return r.Len() > 0 && r.first().HasAttributes(as...)
}

func (r *R) Project(as ...string) *R {
	if !r.HasAttributes(as...) {
		return &R{}
	}
	return r.project(as...)
}

func (r *R) project(as ...string) *R {
	nr := R{}
	for _, t := range *r {
		nr.Add(t.Project(as...))
	}
	return &nr
}

func (r *R) Union(other *R) *R {
	return r.withCompatible(other, r.union)
}

func (r *R) union(other *R) *R {
	nr := R{}
	for _, t := range *r {
		nr.Add(t)
	}
	for _, t := range *other {
		nr.Add(t)
	}
	return &nr
}

func (r *R) Difference(other *R) *R {
	return r.withCompatible(other, r.difference)
}

func (r *R) difference(other *R) *R {
	nr := R{}
	for _, t := range *r {
		if !other.Contains(t) {
			nr.Add(t)
		}
	}
	return &nr
}

func (r *R) IsCompatible(other *R) bool {
	return r.Len() > 0 && other.Len() > 0 && r.isCompatible(other)
}

func (r *R) isCompatible(other *R) bool {
	return r.first().IsCompatible(other.first())
}

func (r *R) first() T { return (*r)[0] }

func (r *R) withCompatible(other *R, f func(other *R) *R) *R {
	if r.IsCompatible(other) {
		return f(other)
	}
	return &R{}
}

func (r *R) NaturalJoin(other *R) *R {
	if r.IsEmpty() || other.IsEmpty() {
		return &R{}
	}
	return r.join(other, r.first().CommonAttributes(other.first())...)
}

func (r *R) join(other *R, on ...string) *R {
	nr := R{}
	r.forEach(other, func(t, t2 T) {
		if !t.IsEqualOn(t2, on...) {
			return
		}
		nt := t.Join(t2)
		if len(on) == 0 {
			nr.add(nt)
		} else {
			nr.Add(nt)
		}
	})
	return &nr
}

func (r *R) forEach(other *R, f func(t, other T)) {
	for _, t := range *r {
		for _, t2 := range *other {
			f(t, t2)
		}
	}
}
