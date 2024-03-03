package table_test

import (
	"testing"

	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func TestProject(t *testing.T) {
	t.Run("Project", func(t *testing.T) {
		u := require.NoError(table.New("name", "age", "city"))
		require.NilError(u.Append(tuple.T{
			"name": "John",
			"age":  42,
			"city": "London"}))
		require.NilError(u.Append(tuple.T{"name": "Jake", "age": 34}))
		actual := u.Project("name", "age")
		expect := require.NoError(table.New("name", "age"))
		expect.Append(
			tuple.T{"name": "John", "age": 42},
			tuple.T{"name": "Jake", "age": 34})
		if !actual.Equals(expect) {
			t.Error("actual and expect should be equal")
		}
	})
}
