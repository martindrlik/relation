package table_test

import (
	"testing"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func TestTable(t *testing.T) {

	t.Run("Equals", func(t *testing.T) {
		u := require.NoError(table.New("name", "age"))
		v := require.NoError(table.New("name", "age"))
		u.Append(tuple.T{"name": "John", "age": 42})
		v.Append(tuple.T{"name": "John", "age": 42})
		if !u.Equals(v) {
			t.Error("u and v should be equal")
		}

		w := require.NoError(table.New("name", "age"))
		w.Append(tuple.T{"name": "Jake", "age": 34})
		if u.Equals(w) {
			t.Error("u and w should not be equal")
		}

		x := require.NoError(table.New("name", "age"))
		x.Append(tuple.T{"name": "Lisa"})
		if u.Equals(x) {
			t.Error("u and x should not be equal")
		}

		z := require.NoError(table.New("name", "age"))
		z.Append(
			tuple.T{"name": "John", "age": 42},
			tuple.T{"name": "Lisa"})
		if u.Equals(z) {
			t.Error("u and z should not be equal")
		}

		q := require.NoError(table.New("name"))
		q.Append(tuple.T{"name": "John"})
		if u.Equals(q) {
			t.Error("u and q should not be equal")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		x := require.NoError(table.New("name", "age"))
		x.Append(tuple.T{"name": "John", "age": 42})
		u := tuple.T{"name": "John", "age": 42}
		v := tuple.T{"name": "Jake", "age": 34}
		w := tuple.T{"city": "London"}
		if !x.Contains(u) {
			t.Error("x should contain u")
		}
		if x.Contains(v) {
			t.Error("x should not contain v")
		}
		if x.Contains(w) {
			t.Error("x should not contain w")
		}
	})

	t.Run("Append", func(t *testing.T) {
		x := require.NoError(table.New("name", "age"))
		u := tuple.T{"name": "John", "age": 42, "city": "London"}
		if err := x.Append(u); err != relation.ErrSchemaMismatch {
			t.Errorf("expected error %v got %v", relation.ErrSchemaMismatch, err)
		}
	})

}
