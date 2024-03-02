package tuple

import (
	"reflect"

	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/schema"
)

type Tuple map[string]any

func (t Tuple) Schema() schema.Schema {
	return schema.NewSchema(maps.Keys(t)...)
}

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
