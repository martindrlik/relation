package maps

func Slice2Set[V comparable](s []V) map[V]struct{} {
	m := make(map[V]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

func Keys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
