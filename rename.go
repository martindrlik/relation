package rex

import "fmt"

func (r *Relation) Rename(rename map[string]string) *Relation {
	s := NewRelation()
	for _, r := range r.relations {
		for _, t := range r.tuples {
			m := map[string]any{}
			for k, v := range t {
				if n, ok := rename[k]; ok {
					if _, ok := t[n]; ok {
						panic(fmt.Sprintf("illegal rename %v to %v", k, n))
					}
					m[n] = v
				} else {
					m[k] = v
				}
			}
			s.InsertTuple(m)
		}
	}
	return s
}
