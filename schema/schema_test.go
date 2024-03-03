package schema_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
)

func TestSchema(t *testing.T) {
	t.Run("Attributes", func(t *testing.T) {
		x := schema.New("name", "age")
		actual := x.Attributes()
		expect := []string{"name", "age"}
		if len(actual) != len(expect) {
			t.Error("actual and expect should have the same length")
		}
		for i := range actual {
			if actual[i] != expect[i] {
				t.Errorf("actual and expect should be equal at index %d", i)
			}
		}
	})
	t.Run("Contains", func(t *testing.T) {
		s := schema.New("name")
		if !s.Contains("name") {
			t.Error("s should contain name")
		}
		if s.Contains("age") {
			t.Error("s should not contain age")
		}
	})

	t.Run("IsEqual", func(t *testing.T) {
		u := schema.New("name", "age")
		v := schema.New("name", "age")
		if !u.IsEqual(v) {
			t.Error("u and v should be equal")
		}

		w := schema.New("name")
		if u.IsEqual(w) {
			t.Error("u and w should not be equal")
		}

		x := schema.New("city", "population")
		if u.IsEqual(x) {
			t.Error("u and x should not be equal")
		}
	})

	t.Run("IsSubset", func(t *testing.T) {
		u := schema.New("name", "age")
		v := schema.New("name", "age")
		if u.IsSubset(v) {
			t.Error("u should be not subset of v")
		}

		w := schema.New("name")
		if !w.IsSubset(u) {
			t.Error("w should be a subset of u")
		}
	})

	t.Run("Intersection", func(t *testing.T) {
		u := schema.New("name", "age")
		v := schema.New("name", "city")
		actual := u.Intersection(v)
		expect := schema.New("name")
		if !actual.IsEqual(expect) {
			t.Error("actual and expect should be equal")
		}
	})
}
