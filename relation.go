package rex

import "github.com/martindrlik/rex/schema"

type R []T

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

func (r *R) Add(t T) *R {
	if !r.IsEmpty() && !schema.IsEqual(t, (*r)[0]) {
		panic("schema mismatch")
	}
	if !r.Contain(t) {
		*r = append(*r, t)
	}
	return r
}

func (r *R) Contain(t T) bool {
	for _, v := range *r {
		if v.Equal(t) {
			return true
		}
	}
	return false
}

func (r *R) IsEmpty() bool {
	return len(*r) == 0
}
