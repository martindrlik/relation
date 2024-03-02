package tuple_test

import (
	"testing"

	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestTuple(t *testing.T) {
	t.Run("Equals", func(t *testing.T) {
		u := tuple.Tuple{"name": "John", "age": 42}
		v := tuple.Tuple{"name": "John", "age": 42}
		w := tuple.Tuple{"name": "Jake", "age": 34}
		x := tuple.Tuple{"city": "London"}
		if !u.Equals(v) {
			t.Error("u and v should be equal")
		}
		if u.Equals(w) {
			t.Error("u and w should not be equal")
		}
		if u.Equals(x) {
			t.Error("u and x should not be equal")
		}
	})

	t.Run("Project", func(t *testing.T) {
		u := tuple.Tuple{"name": "John", "age": 42}
		actual := u.Project(schema.NewSchema("name"))
		expect := tuple.Tuple{"name": "John"}
		if !actual.Equals(expect) {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})
}
