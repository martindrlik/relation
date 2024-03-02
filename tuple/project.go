package tuple

import "github.com/martindrlik/rex/schema"

func (t Tuple) Project(schema schema.Schema) Tuple {
	x := Tuple{}
	for _, a := range schema.Attributes() {
		x[a] = t[a]
	}
	return x
}
