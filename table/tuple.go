package table

import "reflect"

func tupleEqual(a, b map[string]any) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		w, ok := b[k]
		if !ok {
			return false
		}
		if !reflect.DeepEqual(v, w) {
			return false
		}
	}
	return true
}
