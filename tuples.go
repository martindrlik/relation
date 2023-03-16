package rex

import "reflect"

type Tuple []any
type Tuples []Tuple

func (t Tuple) equals(u Tuple) bool {
	for len(t) != len(u) {
		return false
	}
	for i, tv := range t {
		vv := u[i]
		if reflect.TypeOf(tv) != reflect.TypeOf(vv) {
			return false
		}
		switch x := tv.(type) {
		case R:
			if !x.Equals(vv.(R)) {
				return false
			}
		default:
			if !reflect.DeepEqual(tv, vv) {
				return false
			}
		}
	}
	return true
}

func (t Tuples) contains(u Tuple) bool {
	if len(t) > 0 {
		for _, t := range t {
			if t.equals(u) {
				return true
			}
		}
	}
	return false
}

func (t *Tuples) insert(u Tuple) {
	if !t.contains(u) {
		*t = append(*t, u)
	}
}
