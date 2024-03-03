package relation_test

import (
	"testing"

	"github.com/martindrlik/rex/relation"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/tuple"
)

func TestNaturalJoin(t *testing.T) {

	t.Run("NaturalJoin", func(t *testing.T) {
		x := require.NoError(relation.New(tuple.T{"name": "John", "age": 42}))
		y := require.NoError(relation.New(
			tuple.T{"name": "John", "city": "London"},
			tuple.T{"name": "Jake", "city": "Paris"}))
		actual := require.NoError(x.NaturalJoin(y))
		expect := require.NoError(relation.New(
			tuple.T{"name": "John", "age": 42, "city": "London"}))
		if !actual.Equals(expect) {
			t.Error("actual and expect should be equal")
		}

		z := require.NoError(relation.New(tuple.T{"name": "Lisa"}))
		if _, err := x.NaturalJoin(z); err != relation.ErrResultIsEmpty {
			t.Errorf("expected %v, got %v", relation.ErrResultIsEmpty, err)
		}
	})

	t.Run("CartesianProduct", func(t *testing.T) {
		x := require.NoError(relation.New(tuple.T{"name": "John"}))
		y := require.NoError(relation.New(
			tuple.T{"city": "London"},
			tuple.T{"city": "Paris"},
			tuple.T{"city": "New York"}))
		actual := require.NoError(x.NaturalJoin(y))
		expect := require.NoError(relation.New(
			tuple.T{"name": "John", "city": "London"},
			tuple.T{"name": "John", "city": "Paris"},
			tuple.T{"name": "John", "city": "New York"}))
		if !actual.Equals(expect) {
			t.Error("actual and expect should be equal")
		}
	})

}
