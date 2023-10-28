package rex

import "golang.org/x/exp/maps"

func (r *Relation) Projection(a ...string) *Relation {
	pm := map[string]struct{}{}
	for _, a := range a {
		pm[a] = struct{}{}
	}
	nr := NewRelation()
	for k, s := range *r {
		// pick attributes and tuple indexes
		ai := map[string]int{}
		for i, a := range k.split() {
			if _, ok := pm[a]; ok {
				ai[a] = i
			}
		}
		a := maps.Keys(ai)
		i := maps.Values(ai)
		if len(a) == 0 {
			continue
		}
		k := attrs(a).key()

		// pick tuples
		u := make(relation, 0, len(s))
		for _, t := range s {
			pt := make(tuple, 0, len(i))
			for _, i := range i {
				pt = append(pt, t.tuple[i])
			}
			u = append(u, tuplex{pt, meta{t.isPossible}})
		}
		(*nr)[k] = u
	}
	return nr
}
