package relation_test

import (
	"testing"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestRelation(t *testing.T) {

	t.Run("Relation", func(t *testing.T) {
		r := require.Must(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42}))
		if !r.Schema().IsEqual(schema.NewSchema("name", "age")) {
			t.Error("r should have schema name and age")
		}
		if len(r.Tuples()) != 1 {
			t.Error("r should have one tuple")
		}
	})

	t.Run("Equals", func(t *testing.T) {
		r1 := require.Must(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42}))
		r2 := require.Must(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42}))
		r3 := require.Must(relation.NewRelation(tuple.Tuple{"name": "Jake", "age": 34}))
		r4 := require.Must(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42, "city": "London"}))
		r5 := require.Must(relation.NewRelation(
			tuple.Tuple{"name": "John", "age": 42},
			tuple.Tuple{"name": "Jake", "age": 34}))
		if !r1.Equals(r2) {
			t.Error("r1 and r2 should be equal")
		}
		if r1.Equals(r3) {
			t.Error("r1 and r3 should not be equal")
		}
		if r1.Equals(r4) {
			t.Error("r1 and r4 should not be equal (different schema)")
		}
		if r1.Equals(r5) {
			t.Error("r1 and r5 should not be equal (r5 has more tuples)")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		u := tuple.Tuple{"name": "John", "age": 42}
		v := tuple.Tuple{"name": "Jake", "age": 34}
		w := tuple.Tuple{"city": "London"}
		r := require.Must(relation.NewRelation(u, v))
		if !r.Contains(u) {
			t.Error("r should contain u")
		}
		if r.Contains(w) {
			t.Error("r should not contain w")
		}
	})

	t.Run("Append", func(t *testing.T) {
		u := tuple.Tuple{"name": "John", "age": 42}
		v := tuple.Tuple{"name": "Jake", "age": 34}
		w := tuple.Tuple{"city": "London"}
		r := require.Must(relation.NewRelation(u))
		if err := r.Append(v); err != nil {
			t.Error("r should append v")
		}
		if err := r.Append(w); err != relation.ErrSchemaMismatch {
			t.Error("r should not append u due to schema mismatch")
		}

		if _, err := relation.NewRelation(u, w); err != relation.ErrSchemaMismatch {
			t.Error("r should not be created with u and w due to schema mismatch")
		}
	})

	t.Run("Project", func(t *testing.T) {
		r := require.Must(relation.NewRelation(
			tuple.Tuple{"name": "John", "age": 42},
			tuple.Tuple{"name": "John", "age": 34}))
		actual := require.Must(r.Project("name"))
		expect := require.Must(relation.NewRelation(tuple.Tuple{"name": "John"}))
		if !actual.Equals(expect) {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})

	t.Run("Union", func(t *testing.T) {
		r := require.Must(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42}))
		s := require.Must(relation.NewRelation(tuple.Tuple{"name": "Jake", "age": 34}))
		actual := require.Must(r.Union(s))
		expect := require.Must(relation.NewRelation(
			tuple.Tuple{"name": "John", "age": 42},
			tuple.Tuple{"name": "Jake", "age": 34}))
		if !actual.Equals(expect) {
			t.Errorf("expected %v got %v", expect, actual)
		}

		u := require.Must(relation.NewRelation(tuple.Tuple{"name": "John"}))
		if _, err := r.Union(u); err != relation.ErrSchemaMismatch {
			t.Error("r union u should not be possible due to schema mismatch")
		}
	})

}
