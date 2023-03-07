package schema

func (s Schema) IsSubsetOf(t Schema) bool {
	if len(s) > len(t) {
		return false
	}
	for a, st := range s {
		if tt, ok := t[a]; !ok || st != tt {
			return false
		}
	}
	return true
}
