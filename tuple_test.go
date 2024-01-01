package rex

import (
	"sort"
	"testing"
)

func TestTuple(t *testing.T) {
	a := T{"a": 10}
	b := T{"b": 11}
	ab := T{"a": 10, "b": 11}
	ba := T{"a": 10, "b": 11}
	bc := T{"b": 11, "c": 12}
	c := T{"c": 12}
	ax := T{"a": "x"}
	t.Run("IsEqual", func(t *testing.T) {
		if !ab.IsEqual(ba) {
			t.Errorf("expected %v to be equal to %v", ab, ba)
		}
		if ba.IsEqual(bc) {
			t.Errorf("expected %v to be not equal to %v", ba, bc)
		}
	})
	t.Run("IsCompatible", func(t *testing.T) {
		if !ab.IsCompatible(ba) {
			t.Errorf("expected %v to be compatible with %v", ab, ba)
		}
		if ba.IsCompatible(bc) {
			t.Errorf("expected %v to be not compatible with %v", ba, bc)
		}
		if ab.IsCompatible(a) {
			t.Errorf("expected %v to be not compatible with %v", ab, a)
		}
	})
	t.Run("HasAttributes", func(t *testing.T) {
		if a.HasAttributes("a", "b") {
			t.Errorf("expected %v to not have attributes a, b", a)
		}
		if !ab.HasAttributes("a", "b") {
			t.Errorf("expected %v to have attributes a, b", ab)
		}
	})
	t.Run("Project", func(t *testing.T) {
		if !ab.Project("a").IsEqual(a) {
			t.Errorf("expected %v to be equal to %v", ab.Project("a"), a)
		}
	})
	t.Run("CommonAttributes", func(t *testing.T) {
		if !compareSlice(ab.CommonAttributes(ba), "a", "b") {
			t.Errorf("expected %v to be a, b", ab.CommonAttributes(ba))
		}
		common := ab.CommonAttributes(c)
		if len(common) != 0 {
			t.Errorf("expected no common attributes got %v", common)
		}
	})
	t.Run("IsEqualOn", func(t *testing.T) {
		if !a.IsEqualOn(ab, "a") {
			t.Errorf("expected %v to be equal to %v on a", a, ab)
		}
		if !a.IsEqualOn(c) {
			t.Error("if no common attributes IsEqualOn should return true")
		}
		if a.IsEqualOn(ax, "a") {
			t.Errorf("expected %v to be not equal to %v on a", a, ax)
		}
	})
	t.Run("Join", func(t *testing.T) {
		if !a.Join(b).IsEqual(ab) {
			t.Errorf("expected %v to be equal to %v", a.Join(b), ab)
		}
	})
}

func compareSlice[V string](a []V, b ...V) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
