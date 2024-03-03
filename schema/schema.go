package schema

import "github.com/martindrlik/rex/maps"

type Schema struct {
	o []string            // attributes in order
	m map[string]struct{} // attributes as a set
}

func New(attributes ...string) Schema {
	return Schema{
		o: attributes,
		m: maps.Slice2Set(attributes),
	}
}

func (s Schema) Attributes() []string { return s.o }

func (s Schema) Contains(attribute string) bool {
	_, ok := s.m[attribute]
	return ok
}

func (s Schema) IsEqual(other Schema) bool {
	return len(s.o) == len(other.o) && s.isSubsetOrEqual(other)
}

func (s Schema) IsSubset(other Schema) bool {
	return len(s.o) < len(other.o) && s.isSubsetOrEqual(other)
}

func (s Schema) isSubsetOrEqual(other Schema) bool {
	for k := range s.m {
		_, ok := other.m[k]
		if !ok {
			return false
		}
	}
	return true
}

func (s Schema) Intersection(other Schema) Schema {
	x := []string{}
	for k := range s.m {
		if other.Contains(k) {
			x = append(x, k)
		}
	}
	return New(x...)
}
