package schema

import "sort"

func Map(s ...string) map[string]struct{} {
	m := map[string]struct{}{}
	for _, a := range s {
		m[a] = struct{}{}
	}
	return m
}

func Slice[T any](m map[string]T) []string {
	s := make([]string, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	sort.Strings(s)
	return s
}

func IsEqual[T any, V any](t map[string]T, other map[string]V) bool {
	return len(t) == len(other) && isSubsetOrEqual(t, other)
}

func IsSubset[T any, V any](t map[string]T, other map[string]V) bool {
	return len(t) < len(other) && isSubsetOrEqual(t, other)
}

func isSubsetOrEqual[T any, V any](t map[string]T, other map[string]V) bool {
	for k := range t {
		_, ok := other[k]
		if !ok {
			return false
		}
	}
	return true
}
