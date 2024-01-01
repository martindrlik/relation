package rex

import "reflect"

// T represents a tuple.
type T map[string]any

func (t T) IsEqual(other T) bool {
	return len(t) == len(other) && t.isEqual(other)
}

func (t T) isEqual(other T) bool {
	for k, v := range t {
		if !isEqual(v, other[k]) {
			return false
		}
	}
	return true
}

func isEqual(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

func (t T) IsCompatible(other T) bool {
	return len(t) == len(other) && t.isCompatible(other)
}

func (t T) isCompatible(other T) bool {
	for k := range t {
		if _, ok := other[k]; !ok {
			return false
		}
	}
	return true
}

func (t T) HasAttributes(as ...string) bool {
	for _, a := range as {
		if _, ok := t[a]; !ok {
			return false
		}
	}
	return true
}

func (t T) Project(as ...string) T {
	nt := T{}
	for _, a := range as {
		nt[a] = t[a]
	}
	return nt
}

func (t T) CommonAttributes(other T) []string {
	ca := []string{}
	for k := range t {
		if _, ok := other[k]; ok {
			ca = append(ca, k)
		}
	}
	return ca
}

func (t T) IsEqualOn(other T, on ...string) bool {
	for _, k := range on {
		if !isEqual(t[k], other[k]) {
			return false
		}
	}
	return true
}

func (t T) Join(other T) T {
	nt := T{}
	for k, v := range t {
		nt[k] = v
	}
	for k, v := range other {
		nt[k] = v
	}
	return nt
}
