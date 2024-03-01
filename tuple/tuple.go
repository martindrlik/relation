package tuple

import (
	"reflect"
)

type Tuple map[string]any

func (t Tuple) Equals(other Tuple) bool {
	if len(t) != len(other) {
		return false
	}
	for k, v := range t {
		if !reflect.DeepEqual(v, other[k]) {
			return false
		}
	}
	return true
}
