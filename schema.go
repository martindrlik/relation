package rex

type Schema map[string]struct{}

func newSchema(s ...string) Schema {
	m := map[string]struct{}{}
	for _, a := range s {
		m[a] = struct{}{}
	}
	return m
}

func newSchemaTuple(t T) Schema {
	m := map[string]struct{}{}
	for k := range t {
		m[k] = struct{}{}
	}
	return m
}

func (s Schema) IsEqual(other map[string]struct{}) bool {
	return len(s) == len(other) && s.isSubsetOrEqual(other)
}

func (s Schema) IsSubset(other map[string]struct{}) bool {
	return len(s) < len(other) && s.isSubsetOrEqual(other)
}

func (s Schema) isSubsetOrEqual(other map[string]struct{}) bool {
	for k := range s {
		_, ok := other[k]
		if !ok {
			return false
		}
	}
	return true
}
