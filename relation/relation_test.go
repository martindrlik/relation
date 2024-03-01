package relation_test

import (
	"testing"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/schema"
	"github.com/martindrlik/rex/tuple"
)

func TestRelation(t *testing.T) {
	require := func(r *relation.Relation, err error) *relation.Relation {
		if err != nil {
			panic(err)
		}
		return r
	}

	t.Run("Relation", func(t *testing.T) {
		u := tuple.Tuple{"name": "John", "age": 42}
		r := require(relation.NewRelation(u))
		if !r.Schema().IsEqual(schema.NewSchema("name", "age")) {
			t.Error("r should have schema name and age")
		}
		if len(r.Tuples()) != 1 {
			t.Error("r should have one tuple")
		}
	})
	t.Run("Equals", func(t *testing.T) {
		r1 := require(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42}))
		r2 := require(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42}))
		r3 := require(relation.NewRelation(tuple.Tuple{"name": "Jake", "age": 34}))
		r4 := require(relation.NewRelation(tuple.Tuple{"name": "John", "age": 42, "city": "London"}))
		r5 := require(relation.NewRelation(
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
		r := require(relation.NewRelation(u, v))
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
		r := require(relation.NewRelation(u))
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
		u := tuple.Tuple{"name": "John", "age": 42}
		v := tuple.Tuple{"name": "John"}
		r := require(relation.NewRelation(u))
		actual := r.Project(schema.NewSchema("name"))
		expect := require(relation.NewRelation(v))
		if !actual.Equals(expect) {
			t.Errorf("expected %v got %v", expect, actual)
		}
	})
}
