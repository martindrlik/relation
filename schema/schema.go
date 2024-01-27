package schema

func IsEqual(t, other map[string]any) bool {
	return len(t) == len(other) && isSubsetOrEqual(t, other)
}

func IsSubset(t, other map[string]any) bool {
	return len(t) < len(other) && isSubsetOrEqual(t, other)
}

func isSubsetOrEqual(t, other map[string]any) bool {
	for k := range t {
		_, ok := other[k]
		if !ok {
			return false
		}
	}
	return true
}
