package tuple

import "github.com/martindrlik/rex/schema"

func (t T) Project(schema schema.Schema) T {
	x := T{}
	for _, a := range schema.Attributes() {
		x[a] = t[a]
	}
	return x
}
