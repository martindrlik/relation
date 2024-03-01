package tuple

import "github.com/martindrlik/rex/schema"

func (t Tuple) Project(s schema.Schema) Tuple {
	nt := Tuple{}
	for k, v := range t {
		if s.Contains(k) {
			nt[k] = v
		}
	}
	return nt
}
