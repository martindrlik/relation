package tuple

import (
	"reflect"

	"github.com/martindrlik/rex/maps"
	"github.com/martindrlik/rex/schema"
)

type T map[string]any

func (t T) Schema() schema.Schema {
	return schema.New(maps.Keys(t)...)
}

func (t T) Equals(other T) bool {
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
