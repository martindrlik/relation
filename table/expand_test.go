package table_test

import (
	"testing"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func TestExpand(t *testing.T) {

	t.Run("Expand", func(t *testing.T) {
		x := require.NoError(table.New("name", "city"))
		require.NilError(x.Append(tuple.T{"name": "John"}))
		d := require.NoError(table.New("city"))
		require.NilError(d.Append(
			tuple.T{"city": "London"},
			tuple.T{"city": "Paris"},
			tuple.T{"city": "New York"}))
		actual := require.NoError(x.Expand(d))
		expect := require.NoError(table.New("name", "city"))
		require.NilError(expect.Append(
			tuple.T{"name": "John", "city": "London"},
			tuple.T{"name": "John", "city": "Paris"},
			tuple.T{"name": "John", "city": "New York"}))

		if !actual.Equals(expect) {
			t.Errorf("\nexpected\n%vgot\n%v", box.Table(expect), box.Table(actual))
		}
	})

}
