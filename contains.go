package rex

import "reflect"

func (g R) contains(rk string, t []any) bool {
	gr, ok := g[rk]
	if !ok {
		return false
	}
	return gr.contains(t)
}

func (r Relation) contains(t []any) bool {
	for _, rt := range r.tuples {
		if reflect.DeepEqual(rt, t) {
			return true
		}
	}
	return false
}
