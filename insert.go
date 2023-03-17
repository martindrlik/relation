package rex

import (
	"sort"
)

func (r R) Insert(options ...func(*InsertOptions) error) (R, error) {
	opt, err := buildInsertOptions(options...)
	if err != nil {
		return R{}, err
	}

	s := Relation{
		attributes: make([]string, 0, len(opt.src)),
	}
	for k := range opt.src {
		s.attributes = append(s.attributes, k)
	}
	sort.Strings(s.attributes)
	tuple := make(Tuple, 0, len(s.attributes))
	for _, a := range s.attributes {
		e := opt.src[a]
		switch x := e.(type) {
		case map[string]any:
			sr := R{}
			_, err := sr.Insert(Map(x))
			if err != nil {
				return r, err
			}
			tuple = append(tuple, sr)
		default:
			tuple = append(tuple, x)
		}
	}
	s.tuples = append(s.tuples, tuple)
	r.insertRelation(s)
	return r, nil
}

func (r R) insertRelation(s Relation) {
	if len(s.tuples) == 0 {
		return
	}
	rk := s.key()
	gr, ok := r[rk]
	if !ok {
		r[rk] = Relation{attributes: s.attributes}
		gr = r[rk]
	}
	for _, t := range s.tuples {
		gr.tuples.insert(t)
	}
	r[rk] = gr
}
