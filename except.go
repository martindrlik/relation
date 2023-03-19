package rex

import "reflect"

func (r R) Except(s R) R {
	t := R{}
	for k, v := range r {
		if w, ok := s[k]; ok {
			t[k] = v.except(w)
		} else {
			t[k] = v.copy()
		}
	}
	return t
}

func (r Relation) except(s Relation) Relation {
	if r.key() != s.key() {
		panic("incompatible relations")
	}
	t := Relation{}
	t.attributes = make([]string, len(r.attributes))
	copy(t.attributes, r.attributes)
	t.tuples = r.tuples.except(s.tuples)
	return t
}

func (t Tuples) except(u Tuples) Tuples {
	v := Tuples{}
	for _, t := range t {
		for _, u := range u {
			if w, remove := t.except(u); remove {
				continue
			} else if len(w) != len(t) {
				v = append(v, t)
			} else {
				v = append(v, w)
			}
		}
	}
	return v
}

func (t Tuple) except(u Tuple) (v Tuple, remove bool) {
	v = make(Tuple, 0, len(t))
	remove = true
	for i, t := range t {
		u := u[i]
		if reflect.TypeOf(t) != reflect.TypeOf(u) {
			return nil, false
		}
		switch x := t.(type) {
		case R:
			y := x.Except(u.(R))
			if !y.IsEmpty() {
				remove = false
			}
			v = append(v, y)
		default:
			if !reflect.DeepEqual(t, u) {
				return nil, false
			}
			v = append(v, x)
		}
	}
	return
}
