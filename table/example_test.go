package table_test

import (
	"fmt"

	"github.com/martindrlik/rex/box"
	"github.com/martindrlik/rex/require"
	"github.com/martindrlik/rex/table"
	"github.com/martindrlik/rex/tuple"
)

func Example() {

	t := require.Must(table.New("name", "age"))
	t.Append(tuple.T{"name": "John", "age": 42})

	v := require.Must(table.New("name", "age"))
	v.Append(tuple.T{"name": "John", "age": 42})
	v.Append(tuple.T{"name": "Jake"})

	w := require.Must(t.Union(v))
	fmt.Print(box.Table(w.Schema().Attributes(), w.Relations()))

	// Output:
	// ┏━━━━━━┯━━━━━┓
	// ┃ name │ age ┃
	// ┠──────┼─────┨
	// ┃ John │ 42  ┃
	// ┃ Jake │ *   ┃
	// ┗━━━━━━┷━━━━━┛

}
